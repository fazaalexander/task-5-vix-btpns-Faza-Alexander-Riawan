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
	r.POST("/registrasi", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.Run()
}
