package controllers

import (
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllProfile(r *gin.Context) {
	results := models.FindAllProfile()
	r.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Profile",
		Results: results,
	})
}
func DetailusersProfile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.FindProfileByIdUser(id)

	if data.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Profile Found",
			Results: data,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Profile Not Found",
			Results: data,
		})
	}

}
func Createprofile(ctx *gin.Context) {
	account := models.JoinRegist{}
	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile, err := models.CreateProfile(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Response{
			Success: true,
			Message: "Create User success",
			Results: gin.H{
				"id":       profile.Id,
				"fullname": profile.Full_name,
				"email":    account.Email,
			},
		})

}

// newProfile := models.Profile{}

// if err := ctx.ShouldBind(&newProfile); err != nil {
// 	ctx.JSON(http.StatusBadRequest, lib.Response{
// 		Success: false,
// 		Message: "Invalid input data",
// 	})
// 	return
// }

// err := models.CreateProfile(newProfile)
// if err != nil {
// 	ctx.JSON(http.StatusInternalServerError, lib.Response{
// 		Success: false,
// 		Message: "Failed to create Profile",
// 	})
// 	return
// }

// ctx.JSON(http.StatusOK, lib.Response{
// 	Success: true,
// 	Message: "Profile created successfully",
// 	Results: newProfile,
// })
