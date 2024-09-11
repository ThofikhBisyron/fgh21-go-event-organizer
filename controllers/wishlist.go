package controllers

import (
	"net/http"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func DetailWishlist(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	data := models.FindOnewishlist(id)

	if data != nil {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Wishlist Found",
			Results: data,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Wishlist Not Found",
			Results: gin.H{
				"result": data,
			},
		})
	}

}
func CreateWishlist(ctx *gin.Context) {
	user_id := ctx.GetInt("userId")
	if user_id == 0 {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Unauthorized: Invalid or missing user ID",
		})
		return
	}

	var wishlist models.Wishlist
	if err := ctx.ShouldBind(&wishlist); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	createdWishlist, err := models.Createwishlist(wishlist, user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create wishlist: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Wishlist created successfully",
		Results: createdWishlist,
	})
}

func GetUserEventDetails(ctx *gin.Context) {
	userID := ctx.GetInt("userId")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Unauthorized: Invalid or missing user ID",
		})
		return
	}

	eventDetails, err := models.GetEventDetailsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Gagal mengambil detail event: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Detail event berhasil diambil",
		Results: eventDetails,
	})
}
