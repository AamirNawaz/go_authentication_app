package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection

func ConnectMongo() {
	var mongoUrl = "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalf("Mongo Connection failed :%v", err)
	}

	fmt.Println("Mongodb Connected.")
	collection = client.Database("go_auth").Collection("users")

}
