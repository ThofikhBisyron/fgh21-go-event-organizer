package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func useRouterProfile(routersGroup *gin.RouterGroup) {
	routersGroup.GET("/", controllers.ListAllProfile)
	routersGroup.GET("/:id", controllers.DetailusersProfile)
	// routersGroup.POST("/", controllers.Createprofile)
	// routersGroup.PATCH("/:id", controllers.Updateusers)
	// routersGroup.DELETE("/:id", controllers.Deleteusers)

}
