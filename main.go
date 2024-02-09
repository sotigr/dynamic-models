package main

import (
	"fmt"
	"log"
	"main/models"
)

func main() {

	db, err := models.NewDatabase("mongodb://mongo:27017", "myDb")
	if err != nil {
		log.Fatal(err)
	}

	testModel := models.ProductHandle[models.Product](db)
	id, err := testModel.Insert(models.Product{
		Name:  "test",
		Count: 4,
	})

	if err != nil {
		log.Fatal(err)
	}

	outPr, err := testModel.FindId(id)

	fmt.Println(outPr)

}
