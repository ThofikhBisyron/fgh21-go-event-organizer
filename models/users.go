package models

import (
	"fmt"
)

type Data struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

var dataUser = []Data{
	{
		Id:       1,
		Name:     "Fazztrack",
		Email:    "Rais@yahoo.com",
		Password: "12345678",
	},
	{
		Id:       2,
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: "12345678",
	},
}

func FindAllusers() []Data {
	return dataUser
}
func FindOneusers(id int) *Data {
	for _, user := range dataUser {
		if user.Id == id {
			return &user
		}
	}
	return nil
}
func CreateUser(data Data) error {

	for _, user := range dataUser {
		if user.Email == data.Email {
			return fmt.Errorf("email" + data.Email + "is already exist")
		}
	}

	if len(dataUser) > 0 {
		data.Id = dataUser[len(dataUser)-1].Id + 1
	} else {
		data.Id = 1
	}

	dataUser = append(dataUser, data)

	return nil

}

func DeleteUsers(id int) error {
	for i, user := range dataUser {
		if user.Id == id {

			dataUser = append(dataUser[:i], dataUser[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func Updateusers(id int, newData Data) error {
	for i, user := range dataUser {
		if user.Id == id {
			newData.Id = user.Id
			dataUser[i] = newData
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
