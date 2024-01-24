package main

import (
	"backend/models"
	"backend/routers"
	"log"
	"net/http"
)

func main() {
	db := Connect()
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected")
	models.MakeBook(db)

	BookRouter := &routers.BookRouter{DB: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/books/", BookRouter.HandleBooks)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
