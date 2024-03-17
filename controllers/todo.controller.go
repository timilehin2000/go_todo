package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timilehin2000/go_todo/config"
	"github.com/timilehin2000/go_todo/models"
)

var (
	db = config.LoadDb()
)

var todo models.Todo

func AddTodo(w http.ResponseWriter, r *http.Request) {

	json.NewDecoder(r.Body).Decode(&todo)
	statement := `INSERT INTO todos(item, isCompleted) VALUES ($1,$2) returning id, item, isCompleted`

	err := db.QueryRow(statement, todo.Item, todo.IsCompleted).Scan(&todo.Id, &todo.Item, &todo.IsCompleted)

	if err != nil {
		log.Printf("error occurred doing this: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []models.Todo{}

	statement := `SELECT * FROM todos`
	rows, err := db.Query(statement)

	if err != nil {
		log.Fatalf("Error occured %s\n", err.Error())
		fmt.Println(err)
	}

	for rows.Next() {

		err := rows.Scan(&todo.Id, &todo.Item, &todo.IsCompleted)
		if err != nil {
			log.Printf("error occurred doing this: %s", err)
		}

		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["todoId"]

	statement := `SELECT * FROM todos WHERE id = $1`
	rows := db.QueryRow(statement, id)

	err := rows.Scan(&todo.Id, &todo.Item, &todo.IsCompleted)

	if err != nil {
		log.Printf("error occured doing this: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["todoId"]

	statement := `UPDATE todos SET item=$2, isCompleted=$3 where id=$1`
	_, err := db.Exec(statement, id, todo.Item, todo.IsCompleted)

	if err != nil {
		log.Printf("error occurred doing this: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("updated")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["todoId"]

	statement := `DELETE FROM todos WHERE id = $1`
	_, err := db.Exec(statement, id)

	if err != nil {
		log.Printf("error occurred doing this: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}
