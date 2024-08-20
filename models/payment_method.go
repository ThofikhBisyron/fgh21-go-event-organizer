package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Payment_method struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Paymentmethod() []Payment_method {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "payment_methods" order by "id" asc`,
	)
	payment_method, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Payment_method])
	if err != nil {
		fmt.Println(err)
	}
	return payment_method
}
