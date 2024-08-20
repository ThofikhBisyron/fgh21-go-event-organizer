package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
)

type TransactionDetail struct {
	Id             int `json:"id"`
	Transaction_id int `json:"transaction_id" db:"transaction_id"`
	Section_id     int `json:"sectionId" db:"section_id"`
	Ticket_qty     int `json:"ticket_qty" db:"ticket_qty"`
}

func CreateTransactionDetail(data TransactionDetail) TransactionDetail {
	db := lib.Db()
	defer db.Close(context.Background())

	inputSQL := `insert into "transaction_details" (transaction_id, section_id, ticket_qty) values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`
	row := db.QueryRow(context.Background(), inputSQL, data.Transaction_id, data.Section_id, data.Ticket_qty)

	var detail TransactionDetail

	row.Scan(
		&detail.Id,
		&detail.Transaction_id,
		&detail.Section_id,
		&detail.Ticket_qty,
	)
	fmt.Println(row)
	return detail
}
