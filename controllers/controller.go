package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/techrook/gotodo/config"
	"github.com/techrook/gotodo/models"
)

var (
	id        int
	item      string
	completed int
	tmpl      = template.Must(template.ParseFiles("./views/index.html"))
	database  = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &item, &completed)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	_ = tmpl.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

	item := r.FormValue("item")

	_, err := database.Exec(`INSERT INTO todos (item) VALUE (?)`, item)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
