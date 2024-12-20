package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func useRouterCategories(routersGroup *gin.RouterGroup) {
	// routersGroup.Use(middlewares.AddMiddleWares())
	routersGroup.GET("/", controllers.ListAllCategories)
	routersGroup.GET("/:id", controllers.Detailcategories)
	routersGroup.POST("/", controllers.CreateEventCategories)
	routersGroup.PATCH("/:event_id", controllers.Updatecategories)
	routersGroup.DELETE("/:id", controllers.Deletecategories)
	routersGroup.GET("/events/", controllers.FindEvent_Categories)

}
