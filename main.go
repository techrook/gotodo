package main

import (
	"fmt"
	"net/http"
)

var tasks = []string{"code"}

func main() {

	// handle function
	http.HandleFunc("/", serverStarted)
	// handle function : show all tasks
	http.HandleFunc("/tasks", showAllTasks)
	// handle function : add task
	http.HandleFunc("/addtask", addTask)
	// handle function : delete task
	http.HandleFunc("/deletetask", deleteTask)
	// start local server
	http.ListenAndServe(":8080", nil)
}

func serverStarted(w http.ResponseWriter, request *http.Request) {
	// write response
	fmt.Fprintln(w, "Server started successfully @ localhost:8080")
}

func showAllTasks(w http.ResponseWriter, request *http.Request) {
	for _, task := range tasks {
		fmt.Fprintln(w, task)
	}

}

func addTask(w http.ResponseWriter, request *http.Request) {
	task := "eat"
	tasks = append(tasks, task)
	for _, task := range tasks {
		fmt.Fprintln(w, task)
	}
}

func deleteTask(w http.ResponseWriter, request *http.Request) {
	task := "eat"
	for i, t := range tasks {
		if t == task {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	for _, task := range tasks {
		fmt.Fprintln(w, task)
	}
}
