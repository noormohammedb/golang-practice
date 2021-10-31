package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func dbInsertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func dbUpdateOneMovie(movieId string) {

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

func dbDeleteOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	dbRes, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie deleted cound: ", dbRes.DeletedCount)
}

// delete all records form Database

func dbDeleleAllMovie() int64 {

	dbRes, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movies Deleted : ", dbRes.DeletedCount)
	return dbRes.DeletedCount
}

func dbGetAllMovies() []primitive.M {
	dbCursr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for dbCursr.Next(context.Background()) {
		var movie bson.M
		err := dbCursr.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer dbCursr.Close(context.Background())
	return movies
}

// Controller -- file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := dbGetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var userMovie model.Netflix
	json.NewDecoder(r.Body).Decode(&userMovie)
	dbInsertOneMovie(userMovie)

	json.NewEncoder(w).Encode(userMovie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	fmt.Println(params)

	if params["id"] == "" {
		fmt.Println("params id empty : ", params["id"])
		w.Write([]byte("request error"))
		return
	}
	dbUpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode("success")
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	if params["id"] == "" {

		fmt.Println("params id empty : ", params["id"])
		w.Write([]byte("request error"))
		return
	}
	dbDeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("success")
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	delMovCount := dbDeleleAllMovie()

	json.NewEncoder(w).Encode(delMovCount)
}
