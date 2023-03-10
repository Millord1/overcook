package edgedb_manager

import (
	"context"
	"fmt"
	"log"
	"overcook/entities"

	"github.com/edgedb/edgedb-go"
)

func GetAllDishes() []entities.Dish {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var dishes []entities.Dish
	query := "select Dish{title,duration,description,ingredients:{name,@quantity}}"
	err = client.Query(ctx, query, &dishes)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}

	return dishes
}

func AddDish(dishes []entities.Dish) entities.EdgeResponse {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, dish := range dishes {
		// TODO if not exists
		query += fmt.Sprintf("insert Dish{title:='%s',duration:='%d',description:='%s',comment:='%s};",
			dish.Title, dish.Duration, dish.Description, dish.Comment)
	}

	if len(dishes) > 1 {
		// TODO do transaction
	}

	var result entities.EdgeResponse
	err = client.Query(ctx, query, &result)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}

	return result
}
