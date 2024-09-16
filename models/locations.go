package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Location struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Lat   string `json:"lat"`
	Long  string `json:"long"`
	Image string `json:"image"`
}

func FindAlllocation() []Location {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from locations order by "id" asc`,
	)

	location, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Location])
	if err != nil {
		fmt.Println(err)
	}
	return location
}
