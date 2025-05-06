package routes

import (
	"github.com/gorilla/mux"
	"github.com/mostafijurj/go-book-store/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetAllBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/update/{id}", controllers.UpdateBook).Methods("PUT")
}
