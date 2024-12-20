package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Wishlist struct {
	Id       int `json:"id"`
	User_id  int `json:"user_id" db:"user_id"`
	Event_id int `json:"event_id" form:"event_id" db:"event_id"`
}

type EventDetail struct {
	Id          int    `json:"id"`
	EventID     int    `json:"event_id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Image       string `json:"image"`
	Location    *int   `json:"location"`
}

func FindOnewishlist(user_id int) []Wishlist {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "wishlist" where user_id=$1`, user_id,
	)

	wish, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Wishlist])
	if err != nil {
		fmt.Println(err)
	}
	return wish
}

func Createwishlist(data Wishlist, user_id int) (*Wishlist, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	checkQuery := `SELECT COUNT(*) FROM "wishlist" WHERE "event_id" = $1 AND "user_id" = $2`
	var count int
	err := db.QueryRow(context.Background(), checkQuery, data.Event_id, user_id).Scan(&count)

	if err != nil {
		return nil, fmt.Errorf("failed to check existing wishlist: %v", err)
	}

	if count > 0 {
		return nil, fmt.Errorf("event is already in your wishlist")
	}

	query := `INSERT INTO "wishlist" ("event_id", "user_id") VALUES ($1, $2) RETURNING "id", "event_id", "user_id"`
	var detail Wishlist

	err = db.QueryRow(
		context.Background(),
		query,
		data.Event_id,
		user_id,
	).Scan(&detail.Id, &detail.Event_id, &detail.User_id)

	if err != nil {
		return nil, fmt.Errorf("failed to insert wishlist: %v", err)
	}

	return &detail, nil
}

func GetEventDetailsByUserID(userID int) ([]EventDetail, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	query := `
        SELECT 
			w.id as id,
            e.id as event_id, 
            e.tittle, 
            e.description, 
            e.date, 
            e.image,
			e.location
        FROM 
            wishlist w
        JOIN 
            events e ON w.event_id = e.id
        WHERE 
            w.user_id = $1
    `

	rows, err := db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil detail event: %v", err)
	}
	defer rows.Close()

	var eventDetails []EventDetail
	for rows.Next() {
		var detail EventDetail
		if err := rows.Scan(&detail.Id, &detail.EventID, &detail.Tittle, &detail.Description, &detail.Date, &detail.Image, &detail.Location); err != nil {
			return nil, fmt.Errorf("gagal membaca data event: %v", err)
		}
		eventDetails = append(eventDetails, detail)
	}

	return eventDetails, nil
}

func Deletewishlist(id, userID int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	result, err := db.Exec(
		context.Background(),
		`DELETE FROM "wishlist" WHERE id = $1 AND user_id = $2`,
		id, userID,
	)

	if err != nil {
		return fmt.Errorf("failed to delete wishlist: %v", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("wishlist with id %d not found or does not belong to user %d", id, userID)
	}

	return nil
}
func FindOnewishlistbyId(id int) Wishlist {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(), `SELECT * FROM "wishlist" WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error querying database:", err)
	}

	var wishlist Wishlist
	if rows.Next() {
		if err := rows.Scan(&wishlist.Id, &wishlist.User_id, &wishlist.Event_id); err != nil {
			fmt.Println("Error scanning row:", err)
		}
	}

	fmt.Println("Wishlist ditemukan:", wishlist)

	return wishlist
}
