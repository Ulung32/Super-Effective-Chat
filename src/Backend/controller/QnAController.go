package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"Backend/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// Handler
func CreateQnA(c echo.Context) error {
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var table = models.MongoCollection("QnA", client)

	_, err := table.InsertOne(ctx, models.QnA{
		Question: "siapa assisten stima paling baik ",
		Answer:  "yntkts",
	})

	if err != nil {
		fmt.Println("Error Create QnA")
	}
	return c.String(http.StatusOK, "Succesfully created")
}

func GetQnA(c echo.Context) error{
	//mongo connection
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("QnA", client)

	var qnas = make([]models.QnA, 0)

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var qna models.QnA
		if err := cursor.Decode(&qna); err != nil {
			log.Fatal(err)
		}
		qnas = append(qnas, qna)
	}

	// Print the results
	fmt.Println(qnas)

	return c.JSON(http.StatusOK, qnas)

}