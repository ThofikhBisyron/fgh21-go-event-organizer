package controllers

import (
	"fmt"
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

func CreateEventSection(ctx *gin.Context) {
	NewEventSection := []models.Event_sections{}
	fmt.Println(NewEventSection)

	if err := ctx.ShouldBind(&NewEventSection); err != nil {
		fmt.Println("Error in binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid input data",
		})
		return
	}

	for _, section := range NewEventSection {
		err := models.CreateEventsection(section)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "failed to create event sections",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event Sections Created Successfully",
		Results: NewEventSection,
	})

}
