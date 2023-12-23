package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var todos []Todo

func main() {
	r := mux.NewRouter()

	todos = append(todos,
		Todo{ID: 1, Title: "Golang router", Author: "Zoe", Year: "2013"},
		Todo{ID: 2, Title: "Golang Pointer", Author: "Banc", Year: "2013"})

	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", removeTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode((&todo))

	todos = append(todos, todo)

	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])

	i, _ := strconv.Atoi(params["id"])

	for _, todo := range todos {
		if todo.ID == i {
			json.NewEncoder(w).Encode(&todo)
		}
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	json.NewDecoder(r.Body).Decode(&todo)

	for i, item := range todos {
		if item.ID == id {
			todos[i] = todo
		}
	}

	json.NewEncoder(w).Encode(todos)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, item := range todos {
		if item.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(todos)
}
