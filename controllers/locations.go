package controllers

import (
	"net/http"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllLocation(r *gin.Context) {
	result := models.FindAlllocation()
	r.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Locations",
		Results: result,
	})

}
