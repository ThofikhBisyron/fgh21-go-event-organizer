package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllCategories(c *gin.Context) {
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	listEvent, count := models.FindAllCategories(search, page, limit)

	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}
	totalInfo := lib.TotalInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}
	c.JSON(http.StatusOK, lib.Response{
		Success:     true,
		Message:     "success",
		ResultsInfo: totalInfo,
		Results:     listEvent,
	})
}
func Detailcategories(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.FindOnecategories(id)

	if data.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Categories Found",
			Results: data,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Categories Not Found",
			Results: data,
		})
	}

}
func CreateEventCategories(ctx *gin.Context) {
	newCategories := models.Insert_Categories{}

	if err := ctx.ShouldBind(&newCategories); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	err := models.CreateEventcategories(newCategories)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to insert categories",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Categories inserted successfully",
		Results: newCategories,
	})
}

func Updatecategories(ctx *gin.Context) {
	param := ctx.Param("event_id")
	eventID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid event_id parameter",
		})
		return
	}

	var input models.Insert_Categories
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	if input.Category_id == 0 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "category_id is required",
		})
		return
	}

	if err := models.UpdateCategoriesByEventID(eventID, input.Category_id); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to update category for event_id",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Category for event_id successfully updated",
	})
}

func Deletecategories(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataCategories := models.FindOnecategories(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid categories ID",
		})
		return
	}

	err = models.DeleteCategories(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Categories Not Found",
		})
		return

	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Categories deleted successfully",
		Results: dataCategories,
	})

}
func FindEvent_Categories(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	listcategory := models.Findevent_categories(id, page, limit)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List Event Category",
		Results: listcategory,
	})
}
