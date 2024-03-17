package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s\n", err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	credentials := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", credentials)

	if err != nil {
		log.Fatalf("Error connecting to DB %s\n", err.Error())
	} else {
		fmt.Println("Successfully connected!")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// 	_, err = db.Exec(`CREATE DATABASE todo`)

	// 	if err != nil {
	// 		log.Fatalf("Creating DB %s\n", err.Error())
	// 	}

	// 	_, err = db.Exec("Use todo")

	// 	if err != nil {
	// 		log.Fatalf("DB error %s\n", err.Error())
	// 	}

	// 	_, err = db.Exec(`
	// 	CREATE TABLE todos (
	// 		id INT AUTO_INCREMENT,
	// 		item TEXT NOT NULL,
	// 		is_completed BOOLEAN DEFAULT FALSE,
	// 	);
	// `)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	return db
}
