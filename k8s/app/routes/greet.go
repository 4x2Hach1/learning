package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addGreetRoutes(rg *gin.RouterGroup) {
	r := rg.Group("greet")

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello")
	})
	r.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
}
