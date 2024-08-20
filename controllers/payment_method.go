package controllers

import (
	"net/http"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func FindAllPayment(r *gin.Context) {
	results := models.Paymentmethod()
	r.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Payment",
		Results: results,
	})
}
