package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouterpartners(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.ListAllPartners)

}
