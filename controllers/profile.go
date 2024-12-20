package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func ListAllProfileNationalities(r *gin.Context) {
	results := models.FindAllprofilenationalities()
	r.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Nationalities",
		Results: results,
	})
}

func UpdateUserAndProfile(ctx *gin.Context) {

	userId := ctx.GetInt("userId")

	var joinUsers models.JoinUsers
	var profile models.ProfileNoPicture

	if err := ctx.ShouldBind(&joinUsers); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid user data: " + err.Error(),
		})
		return
	}

	if err := ctx.ShouldBind(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid profile data: " + err.Error(),
		})
		return
	}

	err := models.UpdateProfile(userId, joinUsers, profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to update user and profile: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User and profile updated successfully",
		Results: gin.H{
			"profile": profile,
			"user":    joinUsers,
		},
	})
}
func UploadProfileImage(c *gin.Context) {
	id := c.GetInt("userId")
	fmt.Println(id)

	maxFile := 2 * 1024 * 1024
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxFile))

	file, err := c.FormFile("profileImg")
	if err != nil {
		if err.Error() == "http: request body too large" {
			c.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "file size too large, max capacity 500 kb",
			})
			return
		}
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "not file to upload",
		})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "User Not Found",
		})
		return
	}

	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[fileExt] {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "extension file not validate",
		})
		return
	}

	newFile := uuid.New().String() + fileExt

	uploadDir := "./img/profile/"
	if err := c.SaveUploadedFile(file, uploadDir+newFile); err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Upload Failed",
		})
		return
	}

	tes := "http://159.65.11.166:21214/img/profile/" + newFile

	delImgBefore := models.FindProfileByIdUser(id)
	if delImgBefore.Picture != nil {
		fileDel := strings.Split(*delImgBefore.Picture, "8080")[1]
		os.Remove("." + fileDel)
	}

	profile, err := models.UpdateProfileImage(models.Picture{Picture: &tes}, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Upload Failed",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Upload Success",
		Results: gin.H{
			"profile": profile,
		},
	})
}
