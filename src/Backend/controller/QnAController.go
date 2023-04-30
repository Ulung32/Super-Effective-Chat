package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Backend/models"
	"Backend/regex"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// Handler
func CreateQnA(c echo.Context) error {
	// query := c.QueryParam("query")
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

func GetQnA(){
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


	Processor.QnAList = qnas
}

func GetResults(c echo.Context) error{
	GetQnA()
	query := c.QueryParam("query")

	classificationQuery := regex.QueryClassification(query)
	if(classificationQuery == 1){
		index, similarity := Processor.QuerySearch("KMP", query)
		if(similarity > 90){
			// var result, _ = json.Marshal([]Models.Result{{200, asu}})
			var result, _ = json.Marshal(models.Result{200, Processor.QnAList[index].Answer})
			return c.JSON(http.StatusOK, string(result))
		}else{
			var result, _ = json.Marshal(models.Result{404, "Tidak ada pertanyaan yang mirip dalam database kami"})
			return c.JSON(http.StatusOK, string(result))
		}
		
	}else{
		return c.JSON(http.StatusOK, "Belum bisa gan")
	}
	
}