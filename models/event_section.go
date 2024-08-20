package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Event_sections struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Event_id int    `json:"event_id"`
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
