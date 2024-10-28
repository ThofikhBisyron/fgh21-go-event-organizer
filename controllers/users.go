package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllusers(r *gin.Context) {
	search := r.Query("search")
	page, _ := strconv.Atoi(r.Query("page"))
	limit, _ := strconv.Atoi(r.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	if page > 1 {
		limit = (page - 1) * limit
	}
	listUser, count := models.FindAllusers(search, limit, page)

	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}
	totalInfo := lib.TotalInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}
	r.JSON(http.StatusOK, lib.Response{
		Success:     true,
		Message:     "success",
		ResultsInfo: totalInfo,
		Results:     listUser,
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

	id := ctx.GetInt("userId")

	user := models.FindOneusers(id)
	if user.Id == 0 {

		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User with ID not found",
		})
		return
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	models.Updateusers(user.Email, *user.Username, user.Password, id)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User with successfully updated",
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

func UpdatePassword(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	user := models.FindOneusers(id)

	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	var req struct {
		OldPassword string `form:"oldpassword" binding:"required"`
		NewPassword string `form:"password" binding:"required,min=8"`
	}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	fmt.Println("Old password from request:", req.OldPassword)
	fmt.Println("Hashed password from database:", user.Password)

	if !lib.CheckPassword(user.Password, req.OldPassword) {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Old password is incorrect",
		})
		return
	}

	if err := models.Updatepassword(req.NewPassword, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to update password",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Password successfully updated",
	})
}
