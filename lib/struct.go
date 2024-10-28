package lib

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

type Response struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	ResultsInfo TotalInfo   `json:"ResultsInfo,omitempty"`
	Results     interface{} `json:"results,omitempty"`
}

type TotalInfo struct {
	TotalData int `json:"totalData,omitempty"`
	TotalPage int `json:"totalPage,omitempty"`
	Page      int `json:"page,omitempty"`
	Limit     int `json:"limit,omitempty"`
	Next      int `json:"next,omitempty"`
	Prev      int `json:"prev,omitempty"`
}

func CheckPassword(hashedPassword, password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		fmt.Println("Error verifying password:", err)
		return false
	}
	return match
}
