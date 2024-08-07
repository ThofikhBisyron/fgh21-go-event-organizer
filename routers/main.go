package routers

import "github.com/gin-gonic/gin"

func RouterCombine(r *gin.Engine) {
	useRouter(r.Group("/users"))
}
