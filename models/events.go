package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Events struct {
	Id          int     `json:"id"`
	Image       *string `json:"image" form:"image" db:"image"`
	Tittle      *string `json:"tittle" form:"tittle" db:"tittle"`
	Date        *string `json:"date" form:"date" db:"date"`
	Description *string `json:"description" form:"description" db:"description"`
	Location    *int    `json:"location" form:"location" db:"location"`
	Created_by  *int    `json:"created_by" form:"created_by" db:"created_by"`
}

type JoinEvents struct {
	Id          int     `json:"id"`
	Image       *string `json:"image" form:"image" db:"image"`
	Tittle      *string `json:"tittle" form:"tittle" db:"tittle"`
	Date        *string `json:"date" form:"date" db:"date"`
	Description *string `json:"description" form:"description" db:"description"`
	Location    *string `json:"location" form:"location" db:"locations.name"`
	Created_by  *int    `json:"created_by" form:"created_by" db:"created_by"`
}

func FindAllevents(search string) []Events {
	db := lib.Db()
	defer db.Close(context.Background())

	fmt.Println("Search query parameter:", search)

	rows, _ := db.Query(context.Background(),
		`select * from "events" where "tittle" ilike '%' || $1 || '%' order by "id" asc`, search,
	)

	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])
	if err != nil {
		fmt.Println(err)
	}
	return events
}
func FindOneevents(id int) Events {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "events"`,
	)
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])
	if err != nil {
		fmt.Println(err)
	}
	event := Events{}
	for _, v := range events {
		if v.Id == id {
			event = v
		}
	}
	return event
}
func CreateEvents(event Events, id int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `
		INSERT INTO events (image, tittle, date, description, location, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var eventID int
	err := db.QueryRow(
		context.Background(),
		query,
		event.Image,
		event.Tittle,
		event.Date,
		event.Description,
		event.Location,
		id,
	).Scan(&eventID)
	println(err)
	if err != nil {
		return fmt.Errorf("failed to insert event: %w", err)
	}

	return nil
}
func Updateevents(image string, tittle string, date string, description string, location int, created_by int, id string) error {

	db := lib.Db()
	defer db.Close(context.Background())

	dataSql := `update "events" set (image, tittle, date, description, location, created_by) = ($1, $2, $3, $4, $5, $6) where id=$7`

	_, err := db.Exec(context.Background(), dataSql, image, tittle, date, description, location, created_by, id)
	if err != nil {
		return fmt.Errorf("failed to update event: %v", err)
	}

	return nil
}
func DeleteEvent(id int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`delete FROM "events" where id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no event Found")
	}

	return nil
}

func FindeventbyUserId(id int) ([]JoinEvents, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	sql, err := db.Query(context.Background(),
		`
	SELECT 
	events.id,
	events.image,
	events.tittle,
	events."date",
	events.description,
	locations.name,
	events.created_by
	FROM 
	events
	JOIN
	locations ON events.location = locations.id
	WHERE 
	created_by = $1
	`, id)

	if err != nil {
		return []JoinEvents{}, err
	}

	eventId, err := pgx.CollectRows(sql, pgx.RowToStructByPos[JoinEvents])

	if err != nil {
		return []JoinEvents{}, err
	}

	return eventId, err
}
