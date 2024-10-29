package controllers

import (
	"github.com/gin-gonic/gin"
	"go-pizza/initializers"
	"go-pizza/models"
	"net/http"
)

func PizzaGetAll(c *gin.Context) {
	pizza := []models.Pizza{
		models.Pizza{
			Name:  "Пеперони",
			Price: 500,
			Size:  "Large",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"list": pizza,
	})
}

func PizzaCreate(c *gin.Context) {

	var body struct {
		Name  string
		Price uint32
		Size  string
	}

	c.Bind(&body)

	pizza := models.Pizza{
		Name:  body.Name,
		Price: body.Price,
		Size:  body.Size,
	}

	result := initializers.DB.Create(&pizza)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"item": result,
	})
}
