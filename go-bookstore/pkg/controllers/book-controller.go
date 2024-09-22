package controllers

import (
	"encoding/json"
	"go/learn/go-bookstore/pkg/models"
	"go/learn/go-bookstore/pkg/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal("ID should be of int type ", err)
	}
	book, _ := models.GetBookById(id)
	jsonBook, _ := json.Marshal(book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBook)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	NewBook := &models.Book{}
	utils.ParseBody(r, NewBook)
	book := NewBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal("ID should be of int type %s", bookId)
	}
	book := models.DeleteBook(id)
	jsonBook, _ := json.Marshal(book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBook)
}
