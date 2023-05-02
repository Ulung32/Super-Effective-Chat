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

func parseQuery(query string) (string, string) {
    var x, y string
    i := strings.Index(query, "pertanyaan") + len("pertanyaan")
    j := strings.Index(query, "dengan jawaban")
    if i >= 0 && j >= 0 {
        x = strings.TrimSpace(query[i:j])
        y = strings.TrimSpace(query[j+len("dengan jawaban"):])
    }
    return x, y
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
		question, answer := parseQuery(Message.Query)

		qnacoll := models.MongoCollection("QnA", client)

		_, similarity := Processor.QuerySearch(Message.Method, question)

		if(similarity > 90){
			_, err := coll.InsertOne(ctx, models.Chat{
				ID : primitive.NewObjectID(),
				HistoryId: hisId,
				Chat : "Pertanyaan serupa sudah ada di database",
				IsBot: true,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, "")
			}
			return c.JSON(http.StatusOK, "Pertanyaan serupa sudah ada di database")
		}else{
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_, errInsert := qnacoll.InsertOne(ctx, models.QnA{
				ID: primitive.NewObjectID(),
				Question: question,
				Answer:  answer,
			})

			if errInsert != nil {
				fmt.Println("Error Create QnA")
			}
			_, err := coll.InsertOne(ctx, models.Chat{
				ID : primitive.NewObjectID(),
				HistoryId: hisId,
				Chat : "Sukses menambahkan pertanyaan",
				IsBot: true,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, "")
			}
			return c.String(http.StatusOK, "Sukses menambahkan pertanyaan")
		}
		// return c.JSON(http.StatusOK, "Belum bisa gan")
	}else if(classificationQuery == "5"){
		deletedStr := strings.Replace(Message.Query, "hapus pertanyaan ", "", -1)

		qnacoll := models.MongoCollection("QnA", client)
		index, similarity := Processor.QuerySearch(Message.Method, deletedStr)
		if(similarity < 90){
			_, err := coll.InsertOne(ctx, models.Chat{
				ID : primitive.NewObjectID(),
				HistoryId: hisId,
				Chat : "Tidak ada pertanyaan dalam database",
				IsBot: true,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, "")
			}
			return c.JSON(http.StatusOK, "Tidak ada pertanyaan dalam database")
		}else{
			filter := bson.M{"_id": Processor.QnAList[index].ID}
			result, err := qnacoll.DeleteOne(context.Background(), filter)
			if err != nil {
				return c.String(http.StatusInternalServerError, "Error deleting document")
			}

			if result.DeletedCount == 0 {
				return c.String(http.StatusNotFound, "Document not found")
			}
			_, err = coll.InsertOne(ctx, models.Chat{
				ID : primitive.NewObjectID(),
				HistoryId: hisId,
				Chat : "Sukses menghapus pertanyaan",
				IsBot: true,
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, "")
			}
			// Return a success message
			return c.String(http.StatusOK, "Sukses menghapus pertanyaan")
		}
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