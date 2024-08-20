package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Partners struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func FindAllPartners() []Partners {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "partners" order by "id" asc`,
	)

	p, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Partners])
	if err != nil {
		fmt.Println(err)
	}
	return p
}
