package routes

import (
	"github.com/gorilla/mux"
	"github.com/timilehin2000/go_todo/controllers"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todos", controllers.AddTodo).Methods("POST")
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{todoId}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{todoId}", controllers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo/{todoId}", controllers.UpdateTodo).Methods("PUT")

	return router
}
