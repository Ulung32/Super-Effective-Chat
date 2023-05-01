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

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("User", client)
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var userTemp models.User
		if err := cursor.Decode(&userTemp); err != nil {
			log.Fatal(err)
		}
		if(userTemp.UserName == user.UserName){
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username sudah ada, silahkan cari username lain"})
		}
	}



	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var table = models.MongoCollection("User", client)

	_, errInsert := table.InsertOne(ctx, models.User{
		UserName : user.UserName,
		Password:  user.Password,
	})

	if errInsert != nil {
		fmt.Println("Error Create User")
	}
	return c.String(http.StatusOK, "Succesfully created")
}

func GetUser(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	fmt.Println(username)
	fmt.Println(password)
	client := models.MongoConnect()
	defer client.Disconnect(context.TODO())

	coll := models.MongoCollection("User", client)

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var userTemp models.User
		if err := cursor.Decode(&userTemp); err != nil {
			log.Fatal(err)
		}
		fmt.Println(userTemp)
		if(userTemp.UserName == username){
			if(userTemp.Password == password){
				return c.JSON(http.StatusOK, userTemp)
			}else{
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid password"})
			}
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username tidak ditemukan"})
}