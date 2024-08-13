package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Users struct {
	Id       int     `json:"id"`
	Username *string `json:"username" form:"username"`
	Email    string  `json:"email" form:"email" binding:"required,email"`
	Password string  `json:"-" form:"password" binding:"required,min=8"`
}

func FindAllusers() []Users {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "users" order by "id" asc`,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func FindOneusers(id int) Users {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "users"`,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])
	if err != nil {
		fmt.Println(err)
	}
	user := Users{}
	for _, v := range users {
		if v.Id == id {
			user = v
		}
	}
	return user
}
func CreateUser(users Users) error {
	db := lib.Db()
	defer db.Close(context.Background())

	users.Password = lib.Encrypt(users.Password)
	_, err := db.Exec(
		context.Background(),
		`insert into "users" (email, password, username) values ($1, $2, $3)`,
		users.Email, users.Password, users.Username,
	)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}
func FindUserByEmail(email string) Users {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "users"`,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])
	if err != nil {
		fmt.Println(err)
	}
	user := Users{}
	for _, v := range users {
		if v.Email == email {
			user = v
		}
	}
	return user
}

func Updateusers(email string, username string, password string, id string) {
	db := lib.Db()
	defer db.Close(context.Background())

	dataSql := `update "users" set (email, username, password) = ($1, $2, $3) where id=$4`

	db.Exec(context.Background(), dataSql, email, username, password, id)
}
func DeleteUsers(id int) error {
	db := lib.Db()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`delete FROM "users" where id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found")
	}

	return nil
}

// var dataUser = []Data{
// 	{
// 		Id:       1,
// 		Name:     "Fazztrack",
// 		Email:    "Rais@yahoo.com",
// 		Password: "12345678",
// 	},
// 	{
// 		Id:       2,
// 		Name:     "admin",
// 		Email:    "admin@mail.com",
// 		Password: "12345678",
// 	},
// }

// func FindAllusers() []Data {
// 	return dataUser
// }
// func FindOneusers(id int) *Data {
// 	for _, user := range dataUser {
// 		if user.Id == id {
// 			return &user
// 		}
// 	}
// 	return nil
// }
// func CreateUser(data Data) error {

// 	for _, user := range dataUser {
// 		if user.Email == data.Email {
// 			return fmt.Errorf("email" + data.Email + "is already exist")
// 		}
// 	}

// 	if len(dataUser) > 0 {
// 		data.Id = dataUser[len(dataUser)-1].Id + 1
// 	} else {
// 		data.Id = 1
// 	}

// 	dataUser = append(dataUser, data)

// 	return nil

// }

// func DeleteUsers(id int) error {
// 	for i, user := range dataUser {
// 		if user.Id == id {

// 			dataUser = append(dataUser[:i], dataUser[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("user not found")
// }

// func Updateusers(id int, newData Data) error {
// 	for i, user := range dataUser {
// 		if user.Id == id {
// 			newData.Id = user.Id
// 			dataUser[i] = newData
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("user not found")
// }
