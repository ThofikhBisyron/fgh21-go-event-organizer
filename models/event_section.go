package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Event_sections struct {
	Id       int    `json:"id"`
	Name     string `json:"name" db:"name" form:"name"`
	Price    int    `json:"price" db:"price" form:"price"`
	Quantity int    `json:"quantity" db:"quantity" form:"quantity"`
	Event_id int    `json:"event_id" db:"event_id" form:"event_id"`
}

func FindSectionbyeventId(event_id int) ([]Event_sections, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(),
		`select * from "event_sections" where "event_id" = $1 order by "id" asc`, event_id,
	)
	if err != nil {
		fmt.Println(err)
	}
	event_section, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event_sections])
	if err != nil {
		fmt.Println(err)
	}
	return event_section, nil
}

func CreateEventsection(EventSection Event_sections) error {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `INSERT INTO event_sections ("name", "price", "quantity", "event_id") values ($1, $2, $3, $4) RETURNING id`

	EventSectionID := 0
	err := db.QueryRow(
		context.Background(),
		query,
		EventSection.Name,
		EventSection.Price,
		EventSection.Quantity,
		EventSection.Event_id).Scan(&EventSectionID)

	println(err)
	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil
}
