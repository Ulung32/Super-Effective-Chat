package controller

import (
	"Backend/models"
	"Backend/regex"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type message struct{
	Query string
	HistoryId string
	Method string
}
func GetAnswers(c echo.Context) error{
	GetQnA()
	
	Message := new(message)

	if err := c.Bind(Message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	fmt.Println(Message)
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("Chat", client)
	_, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hisId, err := primitive.ObjectIDFromHex(Message.HistoryId)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}	

	_, errInsert := coll.InsertOne(ctx, models.Chat{
		ID : primitive.NewObjectID(),
		HistoryId: hisId,
		Chat : Message.Query,
		IsBot: false,
	})
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, "Error inserting chat")
	}

	classificationQuery := regex.QueryClassification(Message.Query)
	if(classificationQuery == "1"){
		index, similarity := Processor.QuerySearch(strings.ToLower(Message.Method), strings.ToLower(Message.Query))
		if(similarity > 90){
			// var result, _ = json.Marshal([]Models.Result{{200, asu}})
			_, err := coll.InsertOne(ctx, models.Chat{
				ID : primitive.NewObjectID(),
				HistoryId: hisId,
				Chat : Processor.QnAList[index].Answer,
				IsBot: true,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, "Error inserting chat")
			}
			return c.JSON(http.StatusOK, Processor.QnAList[index].Answer)
		}else{
			_, err := coll.InsertOne(ctx, models.Chat{
				ID : primitive.NewObjectID(),
				HistoryId: hisId,
				Chat : "Tidak ada pertanyaan dalam database",
				IsBot: true,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, "Error inserting chat")
			}
			return c.JSON(http.StatusOK, "Tidak ada pertanyaan dalam database")
		}
		
	}else if(classificationQuery == "4"){
		return c.JSON(http.StatusOK, "Belum bisa gan")
	}else if(classificationQuery == "5"){
		return c.JSON(http.StatusOK, "Belum bisa gan")
	}else{
		_, err := coll.InsertOne(ctx, models.Chat{
			ID : primitive.NewObjectID(),
			HistoryId: hisId,
			Chat : classificationQuery,
			IsBot: true,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Error inserting chat")
		}
		return c.JSON(http.StatusOK, classificationQuery)
	}
}