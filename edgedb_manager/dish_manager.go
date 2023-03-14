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

func PushDishes(dishes *[]entities.Dish) []entities.Dish {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, dish := range *dishes {
		dish := dish.FillDefault()
		desc, _ := dish.Description.Get()
		comment, _ := dish.Comment.Get()
		query += fmt.Sprintf("INSERT Dish{title:='%s',duration:=%d,description:='%s',comment:='%s'}unless conflict;",
			dish.Title, dish.Duration, desc, comment)
	}

	var result []entities.Dish

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

func UpdateDish(ingredients *[]entities.Ingredient, steps *[]entities.Step, dish *entities.Dish) []entities.Dish {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	filledDish := dish.FillDefault()
	comment, _ := filledDish.Comment.Get()
	desc, _ := filledDish.Description.Get()

	query := "update Dish filter .title = '%s' set {" + fmt.Sprintf("duration:=%d, description:='%s', comment='%s'",
		dish.Duration, desc, comment)

	for _, ing := range *ingredients {
		ing := ing.FillDefault()
		qtty, _ := ing.Quantity.Get()
		unity, _ := ing.Unity.Get()
		query += fmt.Sprintf("ingredients += (INSERT Ingredient{name:='%s', @quantity:=%d, @unity:='%s'}unless conflict on .name else (select Ingredient));",
			ing.Name, qtty, unity)
	}

	for _, step := range *steps {
		step := step.FillDefault()
		comment, _ := step.Comment.Get()
		query += fmt.Sprintf("steps += (INSERT Step{content:='%s', comment:='%s'}unless conflict on .content else (select Step));",
			step.Content, comment)
	}

	query += "};"

	var result []entities.Dish
	err = client.Query(ctx, query, &result)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func DeleteEntity(entity *entities.EdgeEntity) ([]entities.EdgeEntity, error) {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	query := fmt.Sprintf("delete %s filter .%s = '%s'",
		entities.GetDbName(*entity), entities.GetProperty(*entity), entities.GetValue(*entity))

	var result []entities.EdgeEntity
	err = client.Query(ctx, query, &result)

	return result, err
}

func PushIngredients(ingredients *[]entities.Ingredient) []entities.Ingredient {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, ing := range *ingredients {
		query += fmt.Sprintf("INSERT Ingredient{name:='%s'}unless conflict;", ing.Name)
	}

	var result []entities.Ingredient

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

func PushSteps(steps *[]entities.Step) []entities.Step {

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, step := range *steps {
		step := step.FillDefault()
		comment, _ := step.Comment.Get()
		query += fmt.Sprintf("INSERT Step{content:='%s', comment:='%s'}unless conflict on .content;",
			step.Content, comment)
	}

	var result []entities.Step

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
