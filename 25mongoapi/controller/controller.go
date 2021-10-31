package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/noormohammedb/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// MongoDB helper  -- file

// insert 1 record

func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	fmt.Println("movie update id: ", id)

	filer := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	dbResponse, err := collection.UpdateOne(context.Background(), filer, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update result: ", dbResponse)
}
