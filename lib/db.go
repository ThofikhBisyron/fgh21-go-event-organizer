package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Db() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(),
		"postgresql://postgres:12345678@103.93.58.89:5432/event_organizer?sslmode=disable",
	)

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
