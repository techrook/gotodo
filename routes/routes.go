package routes

import (
	"github.com/gorilla/mux"
	"github.com/techrook/gotodo/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")

	return route
}
