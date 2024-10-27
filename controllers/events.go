package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
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

	if err := ctx.ShouldBind(&newEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	err := models.CreateEvents(newEvent, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create event",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event created successfully",
		Results: newEvent,
	})
}
func UpdateEvents(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	search := ctx.Query("search")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}

	data := models.FindAllevents(search)

	event := models.Events{}
	err = ctx.Bind(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Failed to bind event data: " + err.Error(),
		})
		return
	}

	var existingEvent models.Events
	for _, v := range data {
		if v.Id == id {
			existingEvent = v
			break
		}
	}

	if existingEvent.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Event with ID " + param + " not found",
		})
		return
	}

	err = models.Updateevents(*event.Image, *event.Tittle, *event.Date, *event.Description, *event.Location, *event.Created_by, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to update event: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event with ID " + param + " updated successfully",
		Results: event,
	})
}
func Deleteevent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindOneevents(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid Event ID",
		})
		return
	}

	err = models.DeleteEvent(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Events deleted successfully",
		Results: dataUser,
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
