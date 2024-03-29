package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"overcook/edgedb_manager"
	"overcook/entities"

	"github.com/gin-gonic/gin"
)

func DishesHandler(c *gin.Context) {
	allDishes := edgedb_manager.GetAllDishes()
	result, err := json.Marshal(allDishes)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func AddDish(c *gin.Context) {
	var dish entities.Dish
	if c.ShouldBind(&dish) == nil {
		log.Println(dish.Title)
		log.Println(dish.Duration)
	}

	var dishes []entities.Dish
	dishes = append(dishes, dish)
	edgedb_manager.PushDishes(&dishes)

	c.String(http.StatusOK, "Success")
}

func AddIngredient(c *gin.Context) {
	var ingredient entities.Ingredient
	if c.ShouldBind(&ingredient) == nil {
		log.Println(ingredient.Name)
	}

	var ings []entities.Ingredient
	ings = append(ings, ingredient)
	result := edgedb_manager.PushIngredients(&ings)

	fmt.Println(result)

	c.String(http.StatusOK, "Success")
}

func AddStep(c *gin.Context) {
	var step entities.Step
	if c.ShouldBind(&step) == nil {
		log.Println(step.Content)
	}
	c.String(http.StatusOK, "Success")
}
