package main

import (
	"overcook/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("dish/all", api.DishesHandler)

	router.POST("dish/add", api.AddDish)
	router.POST("ingredient/add", api.AddIngredient)
	router.POST("step/add", api.AddStep)

	router.Run(":8080")
}
