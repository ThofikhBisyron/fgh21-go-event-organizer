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

func CreateEventsection(EventSection *Event_sections) error {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `INSERT INTO event_sections ("name", "price", "quantity", "event_id") values ($1, $2, $3, $4) RETURNING id`

	err := db.QueryRow(
		context.Background(),
		query,
		EventSection.Name,
		EventSection.Price,
		EventSection.Quantity,
		EventSection.Event_id).Scan(&EventSection.Id)

	println(err)
	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil
}

func DeleteSectionsByEventID(eventID int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	_, err := db.Exec(
		context.Background(),
		`DELETE FROM "event_sections" WHERE event_id = $1`,
		eventID,
	)

	if err != nil {
		return fmt.Errorf("failed to delete sections: %v", err)
	}
	return nil
}
func UpsertSection(section *Event_sections) error {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `
		INSERT INTO event_sections (name, price, quantity, event_id)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			price = EXCLUDED.price,
			quantity = EXCLUDED.quantity
		RETURNING id
	`

	if section.Id == 0 {
		query = `
			INSERT INTO event_sections (name, price, quantity, event_id)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`
	}

	err := db.QueryRow(
		context.Background(),
		query,
		section.Name,
		section.Price,
		section.Quantity,
		section.Event_id,
	).Scan(&section.Id)

	if err != nil {
		fmt.Printf("Error during upsert: %v\n", err)
		return fmt.Errorf("failed to upsert section: %v", err)
	}
	fmt.Printf("Upserted section successfully: %+v\n", section)
	return nil
}

func DeleteSectionsNotInIDs(eventID int, ids []int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	if len(ids) == 0 {
		_, err := db.Exec(
			context.Background(),
			`DELETE FROM event_sections WHERE event_id = $1`,
			eventID,
		)
		if err != nil {
			return fmt.Errorf("failed to delete sections: %v", err)
		}
		return nil
	}

	_, err := db.Exec(
		context.Background(),
		`DELETE FROM event_sections WHERE event_id = $1 AND id != ALL($2::int[])`,
		eventID,
		ids,
	)

	if err != nil {
		return fmt.Errorf("failed to delete outdated sections: %v", err)
	}
	return nil
}
func Updatesection(section Event_sections) error {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `UPDATE event_sections 
	          SET name = $1, price = $2, quantity = $3 
	          WHERE id = $4 AND event_id = $5`
	_, err := db.Exec(context.Background(), query,
		section.Name, section.Price, section.Quantity, section.Id, section.Event_id)
	return err
}

func DeleteEventSection(id int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `DELETE FROM event_sections WHERE id = $1`
	_, err := db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete")
	}
	return nil
}

func GetExistingSectionID(eventID int) ([]int, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `SELECT id FROM event_sections WHERE event_id = $1`
	rows, err := db.Query(context.Background(), query, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch existing ID")
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan ID")
		}
		ids = append(ids, id)
	}
	return ids, nil
}
