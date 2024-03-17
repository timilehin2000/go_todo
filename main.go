package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/timilehin2000/go_todo/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	port, existS := os.LookupEnv("PORT")

	if !existS {
		log.Fatal("PORT not set in .env")
	}

	route := routes.Init()
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, route)
}
