package main

import (
	"final-task/controllers"
	"final-task/database"
	"final-task/initializers"
	"final-task/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	database.DbConn()
	database.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/users/register", controllers.Register)
	r.POST("/users/login", controllers.Login)
	r.GET("/users/validate", middlewares.RequireAuth, controllers.Validate)
	r.PUT("/users/:id", controllers.Update)
	r.DELETE("/users/:id", controllers.Delete)
	r.Run()
}
