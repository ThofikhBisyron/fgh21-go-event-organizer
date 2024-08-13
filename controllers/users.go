package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllusers(r *gin.Context) {
	results := models.FindAllusers()
	r.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All User",
		Results: results,
	})
}

func Detailusers(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.FindOneusers(id)

	if data.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Users Found",
			Results: data,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Users Not Found",
			Results: data,
		})
	}

}
func Createusers(ctx *gin.Context) {
	newUser := models.Users{}

	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	err := models.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User created successfully",
		Results: newUser,
	})
}

func Updateusers(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	data := models.FindAllusers()

	user := models.Users{}
	err := ctx.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := models.Users{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "user with id " + param + " not found",
		})
		return
	}
	models.Updateusers(user.Email, *user.Username, user.Password, param)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "user with id " + param + " Edit Success",
		Results: user,
	})
}

func Deleteusers(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindOneusers(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	err = models.DeleteUsers(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id Not Found",
		})
		return

	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User deleted successfully",
		Results: dataUser,
	})

}

// func Detailusers(r *gin.Context) {
// 	id, _ := strconv.Atoi(r.Param("id"))
// 	data := models.FindOneusers(id)

// 	if data.Id != 0 {
// 		r.JSON(http.StatusOK, lib.Response{
// 			Success: true,
// 			Message: "Users OK",
// 			Results: data,
// 		})
// 	} else {
// 		r.JSON(http.StatusNotFound, lib.Response{
// 			Success: false,
// 			Message: "Users Not Found",
// 			Results: data,
// 		})
// 	}
// }

// func Createusers(r *gin.Context) {
// 	newUser := models.Users{}

// 	if err := r.ShouldBind(&newUser); err != nil {
// 		r.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "Invalid input data",
// 		})
// 		return
// 	}

// 	err := models.CreateUser(newUser)
// 	if err != nil {
// 		r.JSON(http.StatusInternalResponseError, lib.Response{
// 			Success: false,
// 			Message: "Failed to create user",
// 		})
// 		return
// 	}

// 	r.JSON(http.StatusOK, lib.Response{
// 		Success: true,
// 		Message: "User created Successsfully",
// 		Results: newUser,
// 	})
// }

// func Updateusers(r *gin.Context) {
// 	param := r.Param("id")
// 	id, err := strconv.Atoi(param)
// 	if err != nil {
// 		r.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "invalid user id",
// 		})
// 		return
// 	}

// 	var user models.Users
// 	if err := r.Bind(&user); err != nil {
// 		r.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "invalid request body",
// 		})
// 		return
// 	}

// 	if err := models.Updateusers(id, user); err != nil {
// 		r.JSON(http.StatusNotFound, lib.Response{
// 			Success: false,
// 			Message: "User Not Found",
// 		})
// 		return
// 	}

// 	r.JSON(http.StatusOK, lib.Response{
// 		Success: true,
// 		Message: "User Successs",
// 		Results: user,
// 	})
// }

// func Deleteusers(r *gin.Context) {
// 	id, err := strconv.Atoi(r.Param("id"))
// 	dataUser := models.FindOneusers(id)

// 	if err != nil {
// 		r.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "Invalid user ID",
// 		})
// 		return
// 	}

// 	err = models.DeleteUsers(id)
// 	if err != nil {
// 		r.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "Id Not Found",
// 		})
// 		return
// 	}

// 	r.JSON(http.StatusOK, lib.Response{
// 		Success: true,
// 		Message: "User deleted Successsfully",
// 		Results: dataUser,
// 	})
// }
