package controllers

import (
	"fmt"
	"net/http"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type Token struct {
	JWToken string `json:"token"`
}

func AuthLogin(ctx *gin.Context) {
	var user models.Users
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	found := models.FindUserByEmail(user.Email)
	fmt.Println(found)
	if found == (models.Users{}) {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong email!",
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)
	fmt.Println(found)
	if isVerified {
		JWToken := lib.GenerateduserIdToken(found.Id)

		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Login success",
			Results: JWToken,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Password",
		})

		return
	}

}
