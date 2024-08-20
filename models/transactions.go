package models

import (
	"context"
	"fmt"
	"time"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	Id                int `json:"id" db:"id"`
	Event_id          int `json:"event_id" form:"event_id" db:"event_id"`
	Payment_method_id int `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
	User_id           int `json:"user_id" form:"user_id" db:"user_id"`
}

type DetailTransaction struct {
	Id                int       `json:"id"`
	Full_name         string    `json:"full_name" form:"full_name" db:"full_name"`
	Event_tittle      string    `json:"event_tittle" form:"event_tittle" db:"tittle"`
	Location_id       *int      `json:"location_id" form:"location_id" db:"location"`
	Date              time.Time `json:"date" form:"date" db:"date"`
	Payment_method_id string    `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
	Section_name      []string  `json:"section_name" form:"section_name" db:"name"`
	Ticket_qty        []int     `json:"ticket_qty" form:"ticket_qty" db:"ticket_qty"`
}

func CreateNewTransactions(data Transaction) Transaction {
	db := lib.Db()
	defer db.Close(context.Background())

	sql := `insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
	row := db.QueryRow(context.Background(), sql, data.Event_id, data.Payment_method_id, data.User_id)

	var results Transaction
	row.Scan(
		&results.Id,
		&results.Event_id,
		&results.Payment_method_id,
		&results.User_id,
	)
	fmt.Println(results)
	return results
}

func AddDetailsTransaction() ([]DetailTransaction, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	sql :=
		`select t.id, p.full_name, e.tittle as "event_tittle", e.location, e.date, pm.name as "payment_method",
        array_agg(es.name) as "section_name", array_agg(td.ticket_qty) as "ticket_qty"
        from "transactions" "t"
        join "users" "u" on u.id = t.user_id
        join "profile" "p" on p.user_id = u.id
        join "events" "e" on e.id = t.event_id
        join "payment_methods" "pm" on pm.id = t.payment_method_id
        join "transaction_details" "td" on td.transaction_id = t.id
        join "event_sections" "es" on es.id = td.section_id
        group by t.id, p.full_name, e.tittle, e.location, e.date, pm.name;`

	send, _ := db.Query(
		context.Background(),
		sql,
	)

	row, err := pgx.CollectRows(send, pgx.RowToStructByPos[DetailTransaction])
	if err != nil {
		fmt.Println(err)
	}
	return row, err
}
func Findtransaction(user_id int) ([]Transaction, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "transactions" where "user_id" = $1`, user_id,
	)
	transaction, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Transaction])
	if err != nil {
		fmt.Println(err)
	}
	return transaction, nil
}
