package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Db() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(),
		"postgresql://postgres:1@157.230.243.8:54321/event_organizer?sslmode=disable",
	)

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
