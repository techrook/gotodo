package routes

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")

	return route
}
