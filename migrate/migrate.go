package main

import (
	"go-pizza/initializers"
	"go-pizza/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	// Модели можно занести в массив и здесь циклом обабатывать, точно также логая ошибку
	err := initializers.DB.AutoMigrate(&models.Pizza{})

	if err != nil {
		log.Fatalf("Failed to migrate model: %v", err)
	}

}
