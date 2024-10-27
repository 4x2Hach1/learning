package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func greet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello!")
}

func greetName(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func addGreetRoutes(rg *gin.RouterGroup) {
	r := rg.Group("greet")
	r.GET("/", greet)
	r.GET("/:name", greetName)
}

