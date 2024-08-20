package controllers

import (
	"net/http"

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
	id := ctx.GetInt("userId")
	data := models.FindProfileByIdUser(id)
	datap := models.FindOneusers(id)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Profile Found",
		Results: gin.H{
			"profile": data,
			"user":    datap,
		},
	})

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
func CreateUserandProfile(ctx *gin.Context) {
	var newUser models.Users
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	userId, err := models.CreateUserprofile(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	tokenUserId := ctx.GetInt("userId")

	if tokenUserId != userId {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	var newProfile models.Profile
	if err := ctx.ShouldBind(&newProfile); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid profile data",
		})
		return
	}

	newProfile.User_id = userId

	err = models.CreateProfileuser(newProfile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create profile",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User and profile created successfully",
		Results: newUser,
	})
}

func ListAllProfileNationalities(r *gin.Context) {
	results := models.FindAllprofilenationalities()
	r.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Nationalities",
		Results: results,
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
