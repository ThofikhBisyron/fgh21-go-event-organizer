package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func useRouterEventSection(routersGroup *gin.RouterGroup) {
	routersGroup.POST("/", controllers.CreateEventSection)
}
