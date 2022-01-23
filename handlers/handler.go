package handlers

import (
	"fmt"
	. "github.com/bilgihankose/golang-bookstore/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func Run(){
	fmt.Println("Running server on :8080")

	router := mux.NewRouter()

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books{id}", removeBook).Methods("DELETE")

	CheckError(http.ListenAndServe(":8080", router))

}


func getAllBooks(w http.ResponseWriter, r *http.Request){
	fmt.Println("getAllBooks is called")
}

func getBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("getBook is called")
}

func addBooks(w http.ResponseWriter, r *http.Request){
	fmt.Println("addBooks is called")
}

func updateBooks(w http.ResponseWriter, r *http.Request){
	fmt.Println("updateBooks is called")
}

func removeBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("removeBook is called")
}