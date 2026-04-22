package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Task model
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task
var idCounter = 1

// Home route
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Task API 🚀")
}

// Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// Create task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	task.ID = idCounter
	idCounter++

	tasks = append(tasks, task)

	json.NewEncoder(w).Encode(task)
}

// Mark task as done
func markDone(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/done/")
	id, _ := strconv.Atoi(idStr)

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Delete task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")
	id, _ := strconv.Atoi(idStr)

	var newTasks []Task

	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		}
	}

	tasks = newTasks
	fmt.Fprintln(w, "Task deleted")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/add", createTask)
	http.HandleFunc("/done/", markDone)
	http.HandleFunc("/delete/", deleteTask)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
