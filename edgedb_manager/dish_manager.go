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

func PushDish(dishes *[]entities.Dish) entities.EdgeResponse {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, dish := range *dishes {
		// TODO if not exists
		query += fmt.Sprintf("INSERT Dish{title:='%s',duration:='%d',description:='%s',comment:='%s}unless conflict on .title;",
			dish.Title, dish.Duration, dish.Description, dish.Comment)
	}

	var result entities.EdgeResponse

	if len(*dishes) > 1 {
		err = client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
			e := tx.Execute(ctx, query)
			return e
		})
	} else {
		err = client.Query(ctx, query, &result)
	}

	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}

	return result
}

func updateDish(ingredients *[]entities.Ingredient, steps *[]entities.Step, dish *entities.Dish) entities.EdgeResponse {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	query := "update Dish filter .title = '%s' set {" + fmt.Sprintf("duration:=%d; description:='%s';", dish.Duration, dish.Description)
	for _, ing := range *ingredients {
		query += fmt.Sprintf("ingredients += (INSERT Ingredient{name:='%s', comment:='%s', @quantity:='%d'}unless conflict on .name else (select Ingredient));",
			ing.Name, ing.Comment, ing.Quantity)
	}
	for _, step := range *steps {
		query += fmt.Sprintf("steps += ()INSERT Step{content:='%s', comment:='%s'}unless conflict on .content else (select Step);",
			step.Content, step.Comment)
	}
	query += "};"

	var result entities.EdgeResponse
	err = client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
		e := tx.Execute(ctx, query)
		return e
	})

	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}

	return result
}

func Delete(entity *entities.EdgeEntity) entities.EdgeResponse {
	result, err := deleteEntity(entities.GetDbName(*entity), entities.GetProperty(*entity), entities.GetValue(*entity))

	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func deleteEntity(entityName string, property string, propertyValue string) (entities.EdgeResponse, error) {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	query := fmt.Sprintf("delete %s filter .%s = '%s'", entityName, property, propertyValue)
	var result entities.EdgeResponse

	err = client.Query(ctx, query, &result)

	return result, err
}

// func DeleteDish(dish *entities.Dish) entities.EdgeResponse {
// 	ctx := context.Background()
// 	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Close()

// 	query := fmt.Sprintf("delete Dish filter .title = '%s'", dish.Title)
// 	var result entities.EdgeResponse

// 	err = client.Query(ctx, query, &result)

// 	if err != nil {
// 		log.Fatalln(err)
// 		fmt.Println(err)
// 	}

// 	return result
// }

func PushIngredient(ingredients *[]entities.Ingredient) entities.EdgeResponse {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, ing := range *ingredients {
		query += fmt.Sprintf("INSERT Ingredient{name:='%s', comment:='%s'}unless conflict on .name;",
			ing.Name, ing.Comment)
	}

	var result entities.EdgeResponse

	if len(*ingredients) > 1 {
		err = client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
			e := tx.Execute(ctx, query)
			return e
		})
	} else {
		err = client.Query(ctx, query, &result)
	}
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}

	return result
}

func PushStep(steps *[]entities.Step) entities.EdgeResponse {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, ing := range *steps {
		query += fmt.Sprintf("INSERT Step{content:='%s', comment:='%s'}unless conflict on .content;",
			ing.Content, ing.Comment)
	}

	var result entities.EdgeResponse

	if len(*steps) > 1 {
		err = client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
			e := tx.Execute(ctx, query)
			return e
		})
	} else {
		err = client.Query(ctx, query, &result)
	}
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}

	return result
}
