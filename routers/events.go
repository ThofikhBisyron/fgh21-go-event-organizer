package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func useRouterEvents(routersGroup *gin.RouterGroup) {
	routersGroup.Use(middlewares.AddMiddleWares())
	routersGroup.GET("/", controllers.ListAllevents)
	routersGroup.GET("/:id", controllers.DetailEvents)
	routersGroup.POST("/", controllers.Createevents)
	routersGroup.PATCH("/:id", controllers.UpdateEvents)
	routersGroup.DELETE("/:id", controllers.Deleteevent)

}
