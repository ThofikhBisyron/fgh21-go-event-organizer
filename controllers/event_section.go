package controllers

import (
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func FindSectionByEventId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := models.FindSectionbyeventId(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Section Not Found",
			Results: data,
		})
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Section Found",
		Results: data,
	})

}
