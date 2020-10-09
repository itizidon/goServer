package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"net/http"
	"io"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-go-driver/bson/primitives"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID primitive.ObjectID 'json:"_id,omitempty" bson:"_id,omitempty"'
	Firstname string 'json:"firstname,omitempty" bson:"firstname,omitempty"'
	Lastname string 'json:"lastname,omitempty" bson:"lastname,omitempty"'

}

var client *mongo.Client

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://lobalhost:27017")
	http.HandleFunc("/person", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println('reached')
	switch r.Method{
	case "POST":
		var person Person = json.NewDecoder(r.Body).Decode(&person)
		collection := client.Database("chef").Collection("people")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		result, _ := collection.InsertOne(ctx, person)
		json.NewEncoder(r).Encode(result)
	}
}
