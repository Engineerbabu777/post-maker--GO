package router

import (
	"postmaking/controllers"

	"github.com/gorilla/mux"
)


func Router() *mux.Router{

	// USING GORILLA MUX!
	router := mux.NewRouter()

	router.HandleFunc("/api/create-post", controllers.HandleCreatePost).Methods("POST");
	router.HandleFunc("/api/get-all-posts", controllers.HandleGetAllPosts).Methods("GET");

	
	return router;
}