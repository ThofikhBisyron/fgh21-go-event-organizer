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
	Password string  `json:"-" form:"password" binding:"required,min=8" db:"password"`
	Role_id  int     `json:"role_id" form:"role_id" db:"role_id"`
}

func FindAllusers(search string, limit int, page int) ([]Users, int) {
	db := lib.Db()
	defer db.Close(context.Background())
	offset := (page - 1) * limit

	sql := `SELECT * FROM "users" where "email" ilike '%' || $1 || '%' limit $2 offset $3`
	rows, _ := db.Query(context.Background(), sql, search, limit, offset)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	}
	result := Total(search)
	return users, result
}
func Totaluser(search string) int {
	db := lib.Db()
	defer db.Close(context.Background())

	sql := `SELECT result(id) as "total" FROM "users" where "email" ilike '%'  $1  '%'`
	rows := db.QueryRow(context.Background(), sql, search)
	var results int
	rows.Scan(
		&results,
	)
	return results
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

func Updateusers(email string, username string, password string, id int) {
	db := lib.Db()
	defer db.Close(context.Background())
	password = lib.Encrypt(password)

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
func Updatepassword(password string, id int) error {
	db := lib.Db()
	defer db.Close(context.Background())
	epassword := lib.Encrypt(password)

	dataSql := `UPDATE "users" SET password = $1 WHERE id = $2`
	_, err := db.Exec(context.Background(), dataSql, epassword, id)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

func CreateUserAndprofile(user Users, profile Profile) error {
	db := lib.Db()
	defer db.Close(context.Background())

	tx, err := db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	user.Password = lib.Encrypt(user.Password)

	err = tx.QueryRow(
		context.Background(),
		`INSERT INTO "users" (email, password, username) 
		 VALUES ($1, $2, $3) RETURNING id`,
		user.Email, user.Password, user.Username,
	).Scan(&user.Id)
	if err != nil {
		tx.Rollback(context.Background())
		return fmt.Errorf("failed to insert user: %v", err)
	}

	profile.User_id = user.Id

	_, err = tx.Exec(
		context.Background(),
		`INSERT INTO "profiles" (full_name, birth_date, gender, phone_number, profession, nationality_id, user_id)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		profile.Full_name, profile.Birth_date, profile.Gender, profile.Phone_number, profile.Profession, profile.Nationality_id, profile.User_id,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return fmt.Errorf("failed to insert profile: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
