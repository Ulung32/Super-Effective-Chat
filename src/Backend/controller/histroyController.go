package controller

import (
	"Backend/models"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type history struct{
	ID primitive.ObjectID
	UserId string
	Name string
}


func CreateHistory(c echo.Context) error{
	History := new(history)
	
	if err := c.Bind(History); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	fmt.Println(History)
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("History", client)
	_, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usId, err := primitive.ObjectIDFromHex(History.UserId)

	// var table = models.MongoCollection("User", client)
	var newHistory = models.History{
		ID: primitive.NewObjectID(),
		UserId:  usId,
		Nama : History.Name,
	}
	_, errInsert := coll.InsertOne(ctx, newHistory)

	if errInsert != nil {
		return c.String(http.StatusInternalServerError, "Server Error")
	}
	return c.JSON(http.StatusOK, newHistory)
}

func DeleteHistory(c echo.Context) error{
	historyid := c.QueryParam("historyID")

	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("History", client)
	
	objID, err := primitive.ObjectIDFromHex(historyid)

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

func GetHistory (c echo.Context) error{
	userID := c.QueryParam("userID")
	fmt.Println(userID)

	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("History", client)

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	
	objID, err := primitive.ObjectIDFromHex(userID)

	fmt.Println(objID)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}	
	
	defer cursor.Close(context.Background())
	var listHistory []models.History
	for cursor.Next(context.Background()) {
		var hisTemp models.History
		if err := cursor.Decode(&hisTemp); err != nil {
			log.Fatal(err)
		}
		if(hisTemp.UserId == objID){
			fmt.Println(hisTemp)
			listHistory = append(listHistory, hisTemp)
		}
	}
	return c.JSON(http.StatusAccepted, listHistory)
}