package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Genre  any    `json:"genre"`
}

type IFace struct {
	DB *sql.DB
}

const (
	suffix = "WHERE id = $1;"
)

func (i *IFace) bookExists(id int, w http.ResponseWriter) (exists bool) {
	row := i.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = $1)", id)
	row.Scan(&exists)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No book with id of " + strconv.Itoa(id)))
	}

	return
}

func checkErr(err error, w http.ResponseWriter) (f bool) {
	f = false

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		f = true
	}
	return
}

func (i *IFace) GetBooks(w http.ResponseWriter, r *http.Request) {
	stmt := `SELECT * FROM books;`
	var books []Book

	rows, err := i.DB.Query(stmt)
	if checkErr(err, w) {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Genre); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	data, err := json.Marshal(books)
	if checkErr(err, w) {
		return
	}
	w.Write(data)
}

func (i *IFace) PostBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if checkErr(err, w) {
		return
	}

	if book.Title == "" || book.Author == "" || book.ISBN == "" {
		http.Error(w, errors.New("Missing one or more of the following values: Title, Author, ISBN").Error(), http.StatusBadRequest)
		return
	}

	var values []interface{}
	stmt := "INSERT INTO books (title, author, isbn"
	values = append(values, book.Title, book.Author, book.ISBN)

	if book.Genre != nil {
		stmt += ", genre) VALUES ($1, $2, $3, $4);"
		values = append(values, book.Genre)
	} else {
		stmt += ") VALUES ($1, $2, $3);"
	}

	_, err = i.DB.Exec(stmt, values...)
	if checkErr(err, w) {
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Resource created successfully"))

}

func (i *IFace) GetBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/books/"))
	if checkErr(err, w) {
		return
	}

	stmt := "SELECT * FROM books " + suffix

	if !i.bookExists(id, w) {
		return
	}

	row := i.DB.QueryRow(stmt, id)
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Genre)
	if checkErr(err, w) {
		return
	}

	data, err := json.Marshal(book)
	if checkErr(err, w) {
		return
	}
	w.Write(data)
}

func (i *IFace) PutBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/books/"))
	if checkErr(err, w) {
		return
	}

	if !i.bookExists(id, w) {
		return
	}

	var book Book
	stmt := "UPDATE books SET "

	err = json.NewDecoder(r.Body).Decode(&book)
	if checkErr(err, w) {
		return
	}

	if book.Title != "" {
		stmt += "title = '" + book.Title + "' "
	}
	if book.Author != "" {
		stmt += "author = '" + book.Author + "' "
	}
	if book.ISBN != "" {
		stmt += "isbn = '" + book.ISBN + "' "
	}
	if book.Genre != nil {
		stmt += "genre = " + fmt.Sprintf("'%v'", book.Genre) + " "
	}

	stmt += suffix
	_, err = i.DB.Exec(stmt, id)
	if checkErr(err, w) {
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully updated book with id of " + strconv.Itoa(id)))
}

func (i *IFace) DelBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/books/"))
	if checkErr(err, w) {
		return
	}

	if !i.bookExists(id, w) {
		return
	}

	stmt := "DELETE FROM books " + suffix
	_, err = i.DB.Exec(stmt, id)
	if checkErr(err, w) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
