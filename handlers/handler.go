package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	. "github.com/bilgihankose/golang-bookstore/models"
	. "github.com/bilgihankose/golang-bookstore/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var books []Book
var db *sql.DB

func Run() {
	fmt.Println("Running server on :8080")

	router := mux.NewRouter()

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books{id}", removeBook).Methods("DELETE")

	CheckError(http.ListenAndServe(":8080", router))

}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books = []Book{}

	rows, err := db.Query("select * from books")
	fmt.Println(err)

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		fmt.Println(err)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)

	defer rows.Close()
}

func getBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	log.Fatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books (title, author, year) values ($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
	log.Fatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	var book Book

	json.NewDecoder(r.Body).Decode(&book)
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 returning id", &book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	log.Fatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)

}

func removeBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("removeBook is called")
}
