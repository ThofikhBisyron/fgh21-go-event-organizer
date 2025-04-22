package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListAllevents(r *gin.Context) {
	search := r.Query("search")

	results := models.FindAllevents(search)
	r.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Events",
		Results: results,
	})
}
func DetailEvents(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.FindOneevents(id)

	if data.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Events Found",
			Results: data,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Events Not Found",
			Results: gin.H{
				"result": data,
			},
		})
	}

}
func Createevents(ctx *gin.Context) {
	var newEvent models.Events
	id := ctx.GetInt("userId")

	tittle := ctx.PostForm("tittle")
	date := ctx.PostForm("date")
	description := ctx.PostForm("description")
	locationStr := ctx.PostForm("location")

	var location *int
	if locationStr != "" {
		loc, err := strconv.Atoi(locationStr)
		if err == nil {
			location = &loc
		}
	}

	newEvent.Tittle = &tittle
	newEvent.Date = &date
	newEvent.Description = &description
	newEvent.Location = location

	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Failed to upload image",
		})
		return
	}
	defer file.Close()

	allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
	fileType := header.Header.Get("Content-Type")
	isValidType := false
	for _, t := range allowedTypes {
		if t == fileType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid file type. Only JPEG, PNG, and GIF are allowed.",
		})
		return
	}

	uniqueID := uuid.New().String()
	fileExt := filepath.Ext(header.Filename)
	newFileName := uniqueID + fileExt

	uploadDir := "./img/events"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create directory",
		})
		return
	}

	filePath := filepath.Join(uploadDir, newFileName)
	out, err := os.Create(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to save image",
		})
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to save image",
		})
		return
	}

	imageURL := "http://localhost:8888/img/events/" + newFileName
	newEvent.Image = &imageURL

	eventID, err := models.CreateEvents(&newEvent, id)
	if err != nil {
		os.Remove(filePath)
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create event",
		})
		return
	}

	newEvent.Id = eventID
	newEvent.Created_by = &id

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event created successfully",
		Results: newEvent,
	})
}
func UpdateEvents(ctx *gin.Context) {
	created_by := ctx.GetInt("userId")
	param := ctx.Param("id")
	eventID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}

	existingEvent, err := models.FindEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to find event: " + err.Error(),
		})
		return
	}

	if existingEvent.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Event with ID " + param + " not found",
		})
		return
	}

	tittle := ctx.PostForm("tittle")
	date := ctx.PostForm("date")
	description := ctx.PostForm("description")
	locationStr := ctx.PostForm("location")

	if tittle == "" || date == "" || description == "" {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Title, date, and description are required fields.",
		})
		return
	}

	location, _ := strconv.Atoi(locationStr)
	var imageURL string
	file, header, err := ctx.Request.FormFile("image")
	if err == nil {
		defer file.Close()

		allowedTypes := map[string]bool{"image/jpeg": true, "image/png": true, "image/gif": true}
		fileType := header.Header.Get("Content-Type")
		if !allowedTypes[fileType] {
			ctx.JSON(http.StatusBadRequest, lib.Response{
				Success: false,
				Message: "Invalid file type. Only JPEG, PNG, and GIF are allowed.",
			})
			return
		}

		uniqueID := uuid.New().String()
		fileExt := filepath.Ext(header.Filename)
		newFileName := uniqueID + fileExt
		uploadDir := "./img/events"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "Failed to create directory",
			})
			return
		}

		filePath := filepath.Join(uploadDir, newFileName)
		out, err := os.Create(filePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "Failed to save image",
			})
			return
		}
		defer out.Close()

		if _, err = io.Copy(out, file); err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Response{
				Success: false,
				Message: "Failed to save image",
			})
			return
		}

		imageURL = "http://localhost:8888/img/events/" + newFileName

		if existingEvent.Image != nil && *existingEvent.Image != "" {
			oldImagePath := "./img/events/" + filepath.Base(*existingEvent.Image)
			if err := os.Remove(oldImagePath); err != nil {
				fmt.Println("Failed to delete old image:", err)
			}
		}
	} else {
		imageURL = *existingEvent.Image
	}

	err = models.Updateevents(
		imageURL,
		tittle,
		date,
		description,
		location,
		created_by,
		param,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to update event: " + err.Error(),
		})
		return
	}

	updatedEvent := models.Events{
		Id:          eventID,
		Tittle:      &tittle,
		Date:        &date,
		Description: &description,
		Location:    &location,
		Image:       &imageURL,
		Created_by:  &created_by,
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event updated successfully",
		Results: updatedEvent,
	})
}
func Deleteevent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid Event ID",
		})
		return
	}

	event := models.FindOneevents(id)

	err = models.DeleteCategoriesByEventID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to delete related categories",
		})
		return
	}

	err = models.DeleteSectionsByEventID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to delete related sections",
		})
		return
	}

	err = models.DeleteEvent(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to delete event",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event and related data deleted successfully",
		Results: event,
	})
}

func FindEventByUserId(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	result, err := models.FindeventbyUserId(id)
	fmt.Println(result)

	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Event Created By User Is Not Found ",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event Created By User Is Found",
		Results: result,
	})
}
