package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"postmaking/helper"
	"postmaking/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// LOADING DOTENV FILE!
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// GETTING VARIABLE!
	uri := os.Getenv("MONGO_URI")
	DB := os.Getenv("DATABASE_NAME")
	CNAME := os.Getenv("COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(uri)

	// CONNECT TO MONGODB!
	client, err := mongo.Connect(context.TODO(), clientOptions)

	helper.ErrorHandler(err)

	println("Connected to MongoDB Successfully!")

	collection = client.Database(DB).Collection(CNAME)
}

// POST CREATION CONTROLLER!
func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPost models.Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	helper.ErrorHandler(err)

	// NOW CREATE POST!
	count, err := collection.InsertOne(context.Background(), newPost)
	helper.ErrorHandler(err)

	var lastInsertedData models.Post
	err = collection.FindOne(context.Background(), bson.M{"_id": count.InsertedID}).Decode(&lastInsertedData)
	helper.ErrorHandler(err)

	//RETURN BACK THE RESPONSE!
	json.NewEncoder(w).Encode(map[string]interface{}{"added": lastInsertedData})

}

// GET ALL POST!
func HandleGetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var allPosts []models.Post
	cur, err := collection.Find(context.TODO(), bson.M{})

	helper.ErrorHandler(err)

	for cur.Next(context.Background()) {

		var post models.Post
		err := cur.Decode(&post)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}

		allPosts = append(allPosts, post)
	}

	json.NewEncoder(w).Encode(allPosts)
}

// UPDATE LIKES!

// UPDATE DISLIKES!

// DELETE POST!

// EDIT POST!
