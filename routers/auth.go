package routers

import (
	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/login", controllers.AuthLogin)
	routerGroup.POST("/register", controllers.Createprofile)

}
