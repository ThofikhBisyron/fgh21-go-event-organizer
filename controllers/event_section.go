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
func UpdateSection(ctx *gin.Context) {
	var sections []models.Event_sections

	eventID, err := strconv.Atoi(ctx.PostForm("event_id"))
	if err != nil || eventID == 0 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Missing or invalid event_id",
		})
		return
	}

	sections = []models.Event_sections{}
	index := 0
	for {

		name := ctx.PostForm(fmt.Sprintf("name_%d", index))
		priceStr := ctx.PostForm(fmt.Sprintf("price_%d", index))
		quantityStr := ctx.PostForm(fmt.Sprintf("quantity_%d", index))

		if name == "" {
			break
		}

		price, _ := strconv.Atoi(priceStr)
		quantity, _ := strconv.Atoi(quantityStr)

		sections = append(sections, models.Event_sections{
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Event_id: eventID,
		})
		index++
	}

	if len(sections) == 0 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "No sections provided",
		})
		return
	}

	existingIDs, err := models.GetExistingSectionID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to fetch existing sections",
		})
		return
	}

	existingIDMap := make(map[int]bool)
	for _, id := range existingIDs {
		existingIDMap[id] = true
	}

	for _, section := range sections {
		fmt.Printf("Processing section: %+v\n", section)

		err := models.UpsertSection(&section)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "Failed to save section",
			})
			return
		}

		fmt.Printf("Section processed successfully: %+v\n", section)
		delete(existingIDMap, section.Id)
	}

	for id := range existingIDMap {
		err := models.DeleteEventSection(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "Failed to delete section",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Sections updated successfully",
	})
}
