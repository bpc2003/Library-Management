package routers

import (
	"backend/controllers"
	"database/sql"
	"errors"
	"net/http"
	"strings"
)

type BookRouter struct {
	DB *sql.DB
}

func (router *BookRouter) HandleBooks(w http.ResponseWriter, r *http.Request) {
	controller := &controllers.IFace{DB: router.DB}
	id := strings.TrimPrefix(r.URL.Path, "/api/books/")

	if id != "" {
		switch r.Method {
		case http.MethodGet:
			controller.GetBook(w, r)
		case http.MethodPut:
			controller.PutBook(w, r)
		case http.MethodDelete:
			controller.DelBook(w, r)
		default:
			http.Error(w, errors.New("Invalid Method").Error(), http.StatusBadRequest)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			controller.GetBooks(w, r)
		case http.MethodPost:
			controller.PostBook(w, r)
		default:
			http.Error(w, errors.New("Invalid Method").Error(), http.StatusBadRequest)
		}
	}
}
