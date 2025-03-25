package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techrook/gotodo/config"
	"github.com/techrook/gotodo/models"
)

var (
	tmpl     = template.Must(template.ParseFiles("./views/index.html"))
	database = config.Database()
)

// Show all tasks
func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT id, item, completed FROM todos`)
	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()

	var todos []models.Todo
	for statement.Next() {
		var todo models.Todo
		err = statement.Scan(&todo.Id, &todo.Item, &todo.Completed)
		if err != nil {
			fmt.Println(err)
		}
		todos = append(todos, todo)
	}

	data := models.View{Todos: todos}
	_ = tmpl.Execute(w, data)
}

// Add a new task
func Add(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	_, err := database.Exec(`INSERT INTO todos (item, completed) VALUES (?, 0)`, item)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Mark a task as completed
func UpdateCompleted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
