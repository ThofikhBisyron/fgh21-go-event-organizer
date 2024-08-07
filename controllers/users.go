package controllers

import (
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllusers(r *gin.Context) {
	results := models.FindAllusers()
	r.JSON(http.StatusOK, lib.Server{
		Succes:  true,
		Message: "List All User",
		Results: results,
	})
}

func Detailusers(r *gin.Context) {
	id, _ := strconv.Atoi(r.Param("id"))
	data := models.FindOneusers(id)

	if data.Id != 0 {
		r.JSON(http.StatusOK, lib.Server{
			Succes:  true,
			Message: "Users OK",
			Results: data,
		})
	} else {
		r.JSON(http.StatusNotFound, lib.Server{
			Succes:  false,
			Message: "Users Not Found",
			Results: data,
		})
	}
}

func Createusers(r *gin.Context) {
	newUser := models.Data{}

	if err := r.ShouldBind(&newUser); err != nil {
		r.JSON(http.StatusBadRequest, lib.Server{
			Succes:  false,
			Message: "Invalid input data",
		})
		return
	}

	err := models.CreateUser(newUser)
	if err != nil {
		r.JSON(http.StatusInternalServerError, lib.Server{
			Succes:  false,
			Message: "Failed to create user",
		})
		return
	}

	r.JSON(http.StatusOK, lib.Server{
		Succes:  true,
		Message: "User created successfully",
		Results: newUser,
	})
}

func Updateusers(r *gin.Context) {
	param := r.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		r.JSON(http.StatusBadRequest, lib.Server{
			Succes:  false,
			Message: "invalid user id",
		})
		return
	}

	var user models.Data
	if err := r.Bind(&user); err != nil {
		r.JSON(http.StatusBadRequest, lib.Server{
			Succes:  false,
			Message: "invalid request body",
		})
		return
	}

	if err := models.Updateusers(id, user); err != nil {
		r.JSON(http.StatusNotFound, lib.Server{
			Succes:  false,
			Message: "User Not Found",
		})
		return
	}

	r.JSON(http.StatusOK, lib.Server{
		Succes:  true,
		Message: "User Success",
		Results: user,
	})
}

func Deleteusers(r *gin.Context) {
	id, err := strconv.Atoi(r.Param("id"))
	dataUser := models.FindOneusers(id)

	if err != nil {
		r.JSON(http.StatusBadRequest, lib.Server{
			Succes:  false,
			Message: "Invalid user ID",
		})
		return
	}

	err = models.DeleteUsers(id)
	if err != nil {
		r.JSON(http.StatusBadRequest, lib.Server{
			Succes:  false,
			Message: "Id Not Found",
		})
		return
	}

	r.JSON(http.StatusOK, lib.Server{
		Succes:  true,
		Message: "User deleted successfully",
		Results: dataUser,
	})
}
