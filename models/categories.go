package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name" form:"name" db:"name"`
}
type Event_Categories struct {
	Id          int    `json:"id"`
	Category_id int    `json:"category_id" db:"id"`
	Name        string `json:"name" db:"name"`
	Image       string `json:"image" db:"image"`
	Date        string `json:"date" db:"date"`
	Tittle      string `json:"tittle" db:"tittle"`
}
type Insert_Categories struct {
	Id          int `json:"id"`
	Event_id    int `json:"event_id" db:"event_id" form:"event_id"`
	Category_id int `json:"category_id" db:"event_id" form:"category_id"`
}

func FindAllCategories(search string, page int, limit int) ([]Categories, int) {
	db := lib.Db()
	defer db.Close(context.Background())
	offset := (page - 1) * limit

	sql := `SELECT * FROM "categories" where "name" ilike '%' || $1 || '%' offset $2 limit $3`
	rows, _ := db.Query(context.Background(), sql, search, offset, limit)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Categories])

	fmt.Println(categories)

	if err != nil {
		fmt.Println(err)
	}
	result := Total(search)
	return categories, result
}
func Total(search string) int {
	db := lib.Db()
	defer db.Close(context.Background())

	sql := `SELECT result(id) as "total" FROM "categories" where "name" ilike '%'  $1  '%'`
	rows := db.QueryRow(context.Background(), sql, search)
	var results int
	rows.Scan(
		&results,
	)
	return results
}
func FindOnecategories(id int) Categories {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "categories"`,
	)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Categories])
	if err != nil {
		fmt.Println(err)
	}
	categorie := Categories{}
	for _, v := range categories {
		if v.Id == id {
			categorie = v
		}
	}
	return categorie
}
func CreateEventcategories(insertCategories Insert_Categories) error {
	db := lib.Db()
	defer db.Close(context.Background())

	_, err := db.Exec(
		context.Background(),
		`insert into event_categories ("event_id", "category_id") values ($1, $2)`,
		insertCategories.Event_id, insertCategories.Category_id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}

func UpdateCategoriesByEventID(eventID int, categoryID int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `UPDATE "event_categories" SET "category_id" = $1 WHERE "event_id" = $2`
	_, err := db.Exec(context.Background(), query, categoryID, eventID)
	return err
}

func DeleteCategories(id int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`delete FROM "categories" where id=$1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no categories found")
	}

	return nil
}
func Findevent_categories(id int, page int, limit int) []Event_Categories {
	db := lib.Db()
	defer db.Close(context.Background())
	offset := (page - 1) * limit

	sql := `SELECT 
			events.id,
			categories.id AS category_id,
			categories.name,
			events.image,
			events."date",
			events.tittle
			FROM
			event_categories
			JOIN 
			events ON event_categories.event_id = events.id
			JOIN
			categories ON event_categories.category_id = categories.id
			WHERE 
			category_id = $1
			ORDER BY event_categories asc
			OFFSET $2 LIMIT $3`

	rows, _ := db.Query(context.Background(), sql, id, offset, limit)
	e_categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event_Categories])

	fmt.Println(e_categories)

	if err != nil {
		fmt.Println(err)
	}
	return e_categories
}

func DeleteCategoriesByEventID(eventID int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	_, err := db.Exec(
		context.Background(),
		`DELETE FROM "event_categories" WHERE event_id = $1`,
		eventID,
	)

	if err != nil {
		return fmt.Errorf("failed to delete categories: %v", err)
	}
	return nil
}
