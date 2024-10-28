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

func CreateEventSection(ctx *gin.Context) {
	names := ctx.PostFormArray("name")
	prices := ctx.PostFormArray("price")
	quantities := ctx.PostFormArray("quantity")
	EventIdStr := ctx.PostForm("event_id")

	if len(names) != len(prices) || len(names) != len(quantities) || EventIdStr == "" {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid input data",
		})
	}

	eventID, err := strconv.Atoi(EventIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid Event ID",
		})
	}

	newSection := []models.Event_sections{}

	for i := range names {
		price, _ := strconv.Atoi(prices[i])
		quantity, _ := strconv.Atoi(quantities[i])

		section := &models.Event_sections{
			Name:     names[i],
			Price:    price,
			Quantity: quantity,
			Event_id: eventID,
		}
		if err := models.CreateEventsection(section); err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "Failed to create event sections",
			})
			return
		}
		newSection = append(newSection, *section)
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event Sections Created Successfully",
		Results: newSection,
	})

}
