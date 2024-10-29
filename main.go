package main

import (
	"github.com/gin-gonic/gin"
	"go-pizza/controllers"
	"go-pizza/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.GET("/pizza", controllers.PizzaGetAll)
	r.POST("/pizza", controllers.PizzaCreate)

	// TODO: Вынести в отдельный сервис
	r.Static("/images", "./static/images")

	r.Run()
}
