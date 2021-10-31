package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://golang-user:golang-password@cluster0.ichzt.mongodb.net/golang-api?retryWrites=true&w=majority"

const dbName = "golang-api"
const collName = "netflix-watchlist"

// IMPORTANT
var collection *mongo.Collection

// Connecting to MongoDB

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("mongodb connection success")
	}

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection Instance Is Ready")
}
