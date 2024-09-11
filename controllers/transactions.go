package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type FormTransactions struct {
	Event_id          int   `json:"event_Id" form:"event_id" db:"event_id"`
	Payment_method_id int   `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
	Section_id        []int `json:"section_id" form:"section_id" db:"section_id"`
	Ticket_qty        []int `json:"ticket_qty" form:"ticket_qty" db:"ticket_qty"`
}

func CreateTransaction(ctx *gin.Context) {
	form := FormTransactions{}

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid input data",
		})
		return
	}

	trx := models.CreateNewTransactions(models.Transaction{
		User_id:           ctx.GetInt("userId"),
		Payment_method_id: form.Payment_method_id,
		Event_id:          form.Event_id,
	})

	for i := range form.Section_id {
		models.CreateTransactionDetail(models.TransactionDetail{
			Section_id:     form.Section_id[i],
			Ticket_qty:     form.Ticket_qty[i],
			Transaction_id: trx.Id,
		})
	}
	details, err := models.AddDetailsTransaction()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Cannot Create Transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Create Transaction",
		"results": details,
	})

}

func ListDetailsTransactions(ctx *gin.Context) {
	details, err := models.AddDetailsTransaction()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch transaction details",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Transaction details retrieved successfully",
		"results": details,
	})
}

func FindTransactionByEventId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := models.Findtransaction(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Transaction Not Found",
			Results: data,
		})
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Transaction Found",
		Results: data,
	})

}
func FindTransactionByUserId(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	fmt.Println((id))
	result, err := models.FindtransactionByuserId(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Transaction User Not Found",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Transaction User Found",
			Results: result,
		})
	}

}
