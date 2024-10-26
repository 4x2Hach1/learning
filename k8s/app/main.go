// https://github.com/gin-gonic/examples/blob/master/group-routes/routes/main.go
package main

import (
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.AddRoutes(r)
	r.Run()
}
