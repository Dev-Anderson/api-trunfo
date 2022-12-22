package routes

import (
	"api-trunfo/controllers"
	"api-trunfo/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.GetAllUser)
			user.POST("/", controllers.CreateUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		trunfo := main.Group("trunfo", middlewares.Auth())
		{
			veiculo := trunfo.Group("veiculo")
			{
				veiculo.GET("/", controllers.GetAllVeiculos)
				veiculo.POST("/", controllers.CreateVeiculos)
				veiculo.GET("/:id", controllers.GetIDVeiculos)
				veiculo.DELETE("/:id", controllers.DeletaVeiculo)
				veiculo.PATCH("/:id", controllers.EditaVeiculo)
			}
		}
	}

	return router
}
