package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/temathc/news-aggregator/pkg/database"
	"github.com/temathc/news-aggregator/pkg/handlers"
)

var conn string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed to load .env")
		log.Fatal(err)
		return
	}
	conn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}

func main() {
	base := database.NewConnect(conn)
	defer base.Close()
	repoDB := database.NewRepoDB(base)

	publications := handlers.NewPublication(repoDB)
	handler := http.NewServeMux()
	handler.HandleFunc("/publications", Logger(publications.ListPublication))

	serv := http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server started...")
	log.Fatal(serv.ListenAndServe())
}

// ------------- Вывод лога --------------

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("Server info: method [%s], connecting from [%v]", r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
