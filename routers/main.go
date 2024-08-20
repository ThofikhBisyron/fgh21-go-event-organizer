package routers

import "github.com/gin-gonic/gin"

func RouterCombine(r *gin.Engine) {
	useRouter(r.Group("/users"))
	AuthRouter(r.Group("/auth"))
	useRouterProfile(r.Group("/profile"))
	useRouterEvents(r.Group("/events"))
	useRouterCategories(r.Group("/categories"))
	useRouterTransactions(r.Group("/transactions"))
	AuthRouterpartners(r.Group("/partners"))

}
