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
func CreateCategories(categories Categories) error {
	db := lib.Db()
	defer db.Close(context.Background())

	_, err := db.Exec(
		context.Background(),
		`insert into "categories" (name) values ($1)`,
		categories.Name,
	)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}

func UpdateCategories(name string, id int) {
	db := lib.Db()
	defer db.Close(context.Background())

	dataSql := `update "categories" set (name) = ($1) where id=$2`

	db.Exec(context.Background(), dataSql, name, id)
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
