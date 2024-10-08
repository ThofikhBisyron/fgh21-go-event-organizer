package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func useRouter(routersGroup *gin.RouterGroup) {
	routersGroup.Use(middlewares.AddMiddleWares())
	routersGroup.GET("/", controllers.ListAllusers)
	routersGroup.GET("/:id", controllers.Detailusers)
	routersGroup.POST("/", controllers.Createusers)
	routersGroup.PATCH("/:id", controllers.Updateusers)
	routersGroup.DELETE("/:id", controllers.Deleteusers)
	routersGroup.PATCH("/password", controllers.UpdatePassword)

}
