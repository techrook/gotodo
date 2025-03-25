package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techrook/gotodo/controllers"
)

func methodOverrideMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if override := r.FormValue("_method"); override != "" {
				r.Method = override
			}
		}
		next.ServeHTTP(w, r)
	})
}

func Init() *mux.Router {
	route := mux.NewRouter()
	route.Use(methodOverrideMiddleware) // Apply middleware

	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/complete/{id}", controllers.UpdateCompleted).Methods("PATCH", "POST") // Allow POST override
	route.HandleFunc("/delete/{id}", controllers.DeleteTask).Methods("DELETE", "POST")       // Allow POST override

	return route
}
