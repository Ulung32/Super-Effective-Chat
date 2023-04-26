package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() *mongo.Client{
	clientOptions := options.Client().ApplyURI("mongodb+srv://Ulung_Stima:QgyYZDJy5.FSnnX@cluster0.b1hig7l.mongodb.net/?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func MongoCollection(coll string, client *mongo.Client) *mongo.Collection {
	return (*mongo.Collection)(client.Database("STIMA").Collection(coll))
}