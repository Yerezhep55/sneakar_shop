package main

import (
	"github.com/gin-gonic/gin"
	"sneaker_shop/config"
	"sneaker_shop/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.RegisterSneakerRoutes(r)
	r.Run(":8080")
}
