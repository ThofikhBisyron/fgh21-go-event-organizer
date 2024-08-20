package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func useRouterTransactions(routersGroup *gin.RouterGroup) {
	routersGroup.Use(middlewares.AddMiddleWares())
	routersGroup.POST("/", controllers.CreateTransaction)
	routersGroup.GET("/", controllers.ListDetailsTransactions)
	routersGroup.GET("/:id", controllers.FindTransactionByEventId)
}
