package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func useRouterLocations(routersGroup *gin.RouterGroup) {
	routersGroup.GET("/", controllers.ListAllLocation)
}
