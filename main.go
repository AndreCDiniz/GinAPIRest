package main

import (
	"github.com/AndreCDiniz/GinAPIRest/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

var pizzas = []models.Pizza{
	{ID: 1, Name: "Toscana", Price: 49.5},
	{ID: 2, Name: "Cheese", Price: 50.9},
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("pizzas", createPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func createPizzas(c *gin.Context) {
	var newPizza models.Pizza
	err := c.ShouldBind(&newPizza)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pizzas = append(pizzas, newPizza)
}

func getPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}
