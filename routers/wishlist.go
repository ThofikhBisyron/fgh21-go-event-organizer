package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func useRouterWishlist(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AddMiddleWares())
	routerGroup.GET("/", controllers.DetailWishlist)
	routerGroup.POST("/", controllers.CreateWishlist)
	routerGroup.GET("/findevent", controllers.GetUserEventDetails)
	routerGroup.DELETE("/:id", controllers.DeleteWishlist)

}
