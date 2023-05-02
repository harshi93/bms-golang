package routes

import (
	"github.com/gorilla/mux"
	"github.com/harshi93/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/listbooks/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/listbook/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/modbook/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/delbook/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
