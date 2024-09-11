package models

import (
	"context"
	"fmt"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Nationalities struct {
	Id   int    `json:"id"`
	Name string `json:"name" form:"name" db:"name"`
}
type JoinRegist struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
	Results  Profile
}

type JoinUsers struct {
	Id       int     `json:"id"`
	Username *string `json:"username" form:"username"`
	Email    string  `json:"email" form:"email"`
}

type Profile struct {
	Id             int     `json:"id"`
	Picture        *string `json:"picture" form:"picture" db:"picture"`
	Full_name      string  `json:"full_name" form:"full_name" db:"full_name"`
	Birth_date     *string `json:"birth_date" form:"birth_date" db:"birth_date"`
	Gender         *int    `json:"gender" form:"gender" db:"gender"`
	Phone_number   *string `json:"phone_number" form:"phone_number" db:"phone_number"`
	Profession     *string `json:"profession" form:"profession" db:"profession"`
	Nationality_id *int    `json:"nationality_id" form:"nationality_id" db:"natinality_id"`
	User_id        int     `json:"user_id" form:"user_id" db:"user_id"`
}

type Picture struct {
	Picture *string `json:"picture" form:"picture" db:"picture"`
	User_id int     `json:"user_id" form:"user_id" db:"user_id"`
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
		`select * from "profile" where "user_id" = $1`, id,
	)
	profile, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[Profile])
	fmt.Println(err)
	if err != nil {

		fmt.Println(err)

	}
	fmt.Println(profile)

	return profile
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

func FindAllprofilenationalities() []Nationalities {
	db := lib.Db()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "nationalities" order by "id" asc`,
	)
	national, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Nationalities])
	if err != nil {
		fmt.Println(err)
	}
	return national
}
func UpdateProfile(userID int, joinUsers JoinUsers, profile Profile) error {
	db := lib.Db()
	defer db.Close(context.Background())

	updateUserQuery := `
		UPDATE "users" 
		SET "username" = $1, "email" = $2
		WHERE "id" = $3
		RETURNING id;
	`
	var updatedUserID int
	err := db.QueryRow(
		context.Background(),
		updateUserQuery,
		joinUsers.Username,
		joinUsers.Email,
		userID,
	).Scan(&updatedUserID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("no user found for the given user ID")
		}
		return fmt.Errorf("failed to update user details: %v", err)
	}

	updateProfileQuery := `
		UPDATE "profile" 
		SET "picture" = $1, 
		    "full_name" = $2, 
		    "birth_date" = $3, 
		    "gender" = $4, 
		    "phone_number" = $5, 
		    "profession" = $6, 
		    "nationality_id" = $7 
		WHERE "user_id" = $8
		RETURNING id;
	`
	var updatedProfileID int
	err = db.QueryRow(
		context.Background(),
		updateProfileQuery,
		profile.Picture,
		profile.Full_name,
		profile.Birth_date,
		profile.Gender,
		profile.Phone_number,
		profile.Profession,
		profile.Nationality_id,
		userID,
	).Scan(&updatedProfileID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("no profile found for the given user ID")
		}
		return fmt.Errorf("failed to update profile: %v", err)
	}

	return nil
}
func UpdateProfileImage(data Picture, id int) (Picture, error) {
	db := lib.Db()
	defer db.Close(context.Background())

	sql := `UPDATE profile SET "picture" = $1 WHERE user_id=$2 returning *`

	row, err := db.Query(context.Background(), sql, data.Picture, id)
	if err != nil {
		return Picture{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Picture])
	if err != nil {
		return Picture{}, nil
	}

	return profile, nil
}
