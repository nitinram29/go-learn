package routes

import (
	"go/learn/go-bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStore = func(route *mux.Router) {
	route.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	route.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	route.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	route.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
