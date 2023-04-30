package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Backend/models"
	"Backend/regex"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// Handler
func CreateQnA(c echo.Context) error {
	GetQnA()
	fmt.Println(Processor.QnAList)
	qna := new(models.QnA)
	
	if err := c.Bind(qna); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	fmt.Println(qna)
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	_, similarity := Processor.QuerySearch("KMP", qna.Question)
	if(similarity > 90){
		// var result, _ = json.Marshal([]Models.Result{{200, asu}})
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Pertanyaan sudah ada di database"})
	}else{
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var table = models.MongoCollection("QnA", client)

		_, err := table.InsertOne(ctx, models.QnA{
			Question: qna.Question,
			Answer:  qna.Answer,
		})

		if err != nil {
			fmt.Println("Error Create QnA")
		}
		return c.String(http.StatusOK, "Succesfully created")
	}

	
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

	// fmt.Println(query)
	// fmt.Println(Processor.QnAList)


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

func DelQnA (c echo.Context) error{
	GetQnA()

	// fmt.Println(Processor.QnAList)
	// for i := 0; i < len(Processor.QnAList); i++{
	// 	fmt.Println(Processor.QnAList[i].ID.Hex())
	// }
	id := c.QueryParam("id")

	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("QnA", client)
	
	objID, err := primitive.ObjectIDFromHex(id)
	fmt.Println(id)
	fmt.Println(objID)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}	
	filter := bson.M{"_id": objID}

	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting document")
	}
	if result.DeletedCount == 0 {
		return c.String(http.StatusNotFound, "Document not found")
	}

	// Return a success message
	return c.String(http.StatusOK, "Document deleted")
}