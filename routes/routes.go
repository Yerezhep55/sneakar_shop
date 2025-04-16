package routes

import (
	"github.com/gin-gonic/gin"
	"sneaker_shop/controllers"
	"sneaker_shop/middleware"
)

func RegisterSneakerRoutes(r *gin.Engine) {

	sneakerRoutes := r.Group("/sneakers")
	{

		sneakerRoutes.GET("/", controllers.GetSneakers)
		sneakerRoutes.GET("/:id", controllers.GetSneakerByID)

		authorized := sneakerRoutes.Group("/")
		authorized.Use(middleware.JWTAuthMiddleware())
		{
			authorized.POST("/", controllers.CreateSneaker)

			authorized.PUT("/:id", controllers.UpdateSneaker)

			authorized.DELETE("/:id", controllers.DeleteSneaker)
		}
	}
}
