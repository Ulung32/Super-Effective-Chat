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

func CreateHistory(c echo.Context) error{
	history := new(models.History)
	
	if err := c.Bind(history); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	fmt.Println(history)
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("History", client)
	_, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// var table = models.MongoCollection("User", client)

	_, errInsert := coll.InsertOne(ctx, models.History{
		UserId:  history.UserId,
		Nama : history.Nama,
	})

	if errInsert != nil {
		fmt.Println("Error Create User")
	}
	return c.String(http.StatusOK, "Succesfully created")
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
	for cursor.Next(context.Background()) {
		var hisTemp models.History
		if err := cursor.Decode(&hisTemp); err != nil {
			// fmt.Println(hisTemp.UserId)
			log.Fatal(err)
		}
		fmt.Println(hisTemp)
		if(hisTemp.UserId == objID){
			return c.JSON(http.StatusAccepted, hisTemp)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "history tidak ditemukan"})
}