package models

import (
	"context"
	"fmt"
	"time"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Events struct {
	Id          int       `json:"id"`
	Image       *string   `json:"image" form:"image" db:"image"`
	Tittle      *string   `json:"tittle" form:"tittle" db:"tittle"`
	Date        time.Time `json:"date" form:"date" db:"date"`
	Description *string   `json:"description" form:"description" db:"description"`
	Location    *int      `json:"location" form:"location" db:"location"`
	Created_by  *int      `json:"created_by" form:"created_by" db:"created_by"`
}

func FindAllevents() []Events {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "events" order by "id" asc`,
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

	_, err := db.Exec(
		context.Background(),
		`insert into "events" (image, tittle, date, description, location, created_by) values ($1, $2, $3, $4, $5, $6)`,
		event.Image, event.Tittle, event.Date, event.Description, event.Location, id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}
func Updateevents(image string, tittle string, date time.Time, description string, location int, created_by int, id string) error {

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
