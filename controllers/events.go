package controllers

import (
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllevents(r *gin.Context) {
	results := models.FindAllevents()
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
	newEvents := models.Events{}
	id, _ := ctx.Get("userId")

	if err := ctx.ShouldBind(&newEvents); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	err := models.CreateEvents(newEvents, id.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create Profile",
		})
		return
	}
	newEvents.Created_by = id.(*int)
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event created successfully",
		Results: newEvents,
	})
}
func UpdateEvents(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}

	data := models.FindAllevents()

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

	err = models.Updateevents(*event.Image, *event.Tittle, event.Date, *event.Description, *event.Location, *event.Created_by, param)
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
