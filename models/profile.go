package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type JoinRegist struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
	Results  Profile
}

type Profile struct {
	Id             int     `json:"id"`
	Picture        *string `json:"picture" form:"picture"`
	Full_name      string  `json:"full_name" form:"full_name"`
	Birth_date     *string `json:"birth_date" form:"birth_date"`
	Gender         *int    `json:"gender" form:"gender"`
	Phone_number   *string `json:"phone_number" form:"phone_number"`
	Profession     *string `json:"profession" form:"profession"`
	Nationality_id *int    `json:"nationality_id" form:"nationality_id"`
	User_id        int     `json:"user_id" form:"user_id"`
}

func FindAllProfile() []Profile {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "profile" order by "id" asc`,
	)
	profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])
	if err != nil {
		fmt.Println(err)
	}
	return profile
}
func FindProfileByIdUser(id int) Profile {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "profile"`,
	)
	profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])
	if err != nil {
		fmt.Println(err)
	}
	dataprofile := Profile{}
	for _, v := range profile {
		if v.Id == id {
			dataprofile = v
		}
	}
	return dataprofile
}

func CreateProfile(joinRegist JoinRegist) (*Profile, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	joinRegist.Password = lib.Encrypt(joinRegist.Password)

	var userId int
	err := db.QueryRow(
		context.Background(),
		`INSERT INTO "users" ("email", "password") VALUES ($1, $2) RETURNING "id"`,
		joinRegist.Email, joinRegist.Password,
	).Scan(&userId)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into users table: %v", err)
	}
	fmt.Println("-----")
	fmt.Println(err)

	profile := Profile{
		Full_name: joinRegist.Results.Full_name,
		User_id:   userId,
	}
	err = db.QueryRow(
		context.Background(),
		`INSERT INTO "profile" ("picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id", "user_id") 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, picture, full_name, birth_date, gender, phone_number, profession, nationality_id, user_id`,
		joinRegist.Results.Picture, joinRegist.Results.Full_name, joinRegist.Results.Birth_date, joinRegist.Results.Gender,
		joinRegist.Results.Phone_number, joinRegist.Results.Profession, joinRegist.Results.Nationality_id, userId,
	).Scan(
		&profile.Id, &profile.Picture, &profile.Full_name, &profile.Birth_date,
		&profile.Gender, &profile.Phone_number, &profile.Profession, &profile.Nationality_id, &profile.User_id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert into profile table: %v", err)
	}

	return &profile, nil
}
