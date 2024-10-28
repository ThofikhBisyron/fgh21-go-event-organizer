package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func useRouterProfile(routersGroup *gin.RouterGroup) {
	routersGroup.GET("/national", controllers.ListAllProfileNationalities)
	routersGroup.Use(middlewares.AddMiddleWares())
	routersGroup.GET("/", controllers.DetailusersProfile)
	routersGroup.PATCH("/update", controllers.UpdateUserAndProfile)
	routersGroup.PATCH("/", controllers.UploadProfileImage)

}
