package routes

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	addGreetRoutes(v1)
}
