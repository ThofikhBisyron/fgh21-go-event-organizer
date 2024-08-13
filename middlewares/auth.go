package middlewares

import (
	"net/http"

	"github.com/ThofikhBisyron/fgh21-react-go-event-organizer/lib"
	"github.com/gin-gonic/gin"
)

func tokenfailed(ctx *gin.Context) {
	if e := recover(); e != nil {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Unnauthorized",
		})
		ctx.Abort()
	}
}
func AddMiddleWares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer tokenfailed(ctx)
		token := ctx.GetHeader("Authorization")[7:]
		isValidated, userId := lib.ValidateToken(token)
		if isValidated {
			ctx.Set("userId", userId)
			ctx.Next()
		} else {
			panic("Error: Token Invalid")
		}
	}
}
