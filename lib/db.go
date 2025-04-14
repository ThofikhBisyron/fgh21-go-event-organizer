package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Db() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(),
		"postgresql://postgres:1@31.56.56.74:54324/event_organizer?sslmode=disable",
	)

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
