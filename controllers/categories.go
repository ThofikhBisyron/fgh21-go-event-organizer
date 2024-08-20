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

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	if page > 1 {
		page = (page - 1) * limit
	}
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
func Createcategories(ctx *gin.Context) {
	newCategories := models.Categories{}

	if err := ctx.ShouldBind(&newCategories); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	err := models.CreateCategories(newCategories)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Categories created successfully",
		Results: newCategories,
	})
}

func Updatecategories(ctx *gin.Context) {

	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	categorie := models.FindOnecategories(id)
	if categorie.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Categories with ID " + param + " not found",
		})
		return
	}

	if err := ctx.ShouldBind(&categorie); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	models.UpdateCategories(categorie.Name, id)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "categories with ID " + param + " successfully updated",
		Results: categorie,
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
