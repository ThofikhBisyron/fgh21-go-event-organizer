package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Db() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(),
		"postgresql://postgres:12345678@172.17.0.2/event_organizer?sslmode=disable",
	)

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
