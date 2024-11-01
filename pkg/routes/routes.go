package routes

import (
	"github.com/crud/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBookById).Methods("DELETE")

}
