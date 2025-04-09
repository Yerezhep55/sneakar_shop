package routes

import (
	"github.com/gin-gonic/gin"
	"sneaker_shop/controllers"
)

func RegisterSneakerRoutes(r *gin.Engine) {
	sneakerRoutes := r.Group("/sneakers")
	{
		sneakerRoutes.GET("/", controllers.GetSneakers)
		sneakerRoutes.GET("/:id", controllers.GetSneakerByID)
		sneakerRoutes.POST("/", controllers.CreateSneaker)
		sneakerRoutes.PUT("/:id", controllers.UpdateSneaker)
		sneakerRoutes.DELETE("/:id", controllers.DeleteSneaker)
	}
}
