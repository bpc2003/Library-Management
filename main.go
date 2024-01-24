package main

import (
	"backend/controllers"
	"backend/models"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := Connect()
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB Connected")
	models.MakeBook(db)

	books := &controllers.IFace{DB: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/books/", books.HandleBooks)

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
