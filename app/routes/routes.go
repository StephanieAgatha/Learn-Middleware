package routes

import (
	"github.com/gin-gonic/gin"
	"learn-middleware-example/app/controllers"
	"learn-middleware-example/app/middlewares"
)

func SetupNewRoutes() *gin.Engine {
	r := gin.Default()

	//create group for auth and welcome handler
	authRoute := r.Group("/auth")
	authRoute.POST("/login", controllers.LoginHandler)
	authRoute.POST("/register", controllers.RegisterHandler)

	//define middleware in here
	validRoute := r.Group("/v1")
	validRoute.Use(middlewares.AuthMiddleware())
	// handler
	validRoute.GET("/welcome", controllers.WelcomeHandler)

	r.Run(":3000")
	return r
}
