package routers

import (
	"backend/controllers"
	"backend/middleware"
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
			middleware.Auth(controller.PutBook).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Auth(controller.DelBook).ServeHTTP(w, r)
		default:
			http.Error(w, errors.New("Invalid Method").Error(), http.StatusBadRequest)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			controller.GetBooks(w, r)
		case http.MethodPost:
			middleware.Auth(controller.PostBook).ServeHTTP(w, r)
		default:
			http.Error(w, errors.New("Invalid Method").Error(), http.StatusBadRequest)
		}
	}
}
