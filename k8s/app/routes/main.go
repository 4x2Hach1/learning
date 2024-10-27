package routes

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.Engine) {
	addGreetRoutes(r.Group("v1"))
}
