package handlers

import (
	"appdynamics.com/golang/http-mux/books/services"
	"github.com/gorilla/mux"
)

func NewBookHandler(svc services.BookService) BookHandler {
	handler := BookHandler{service: svc}
	handler.Router = mux.NewRouter()
	handler.setupRoutes()
	return handler
}

func (h BookHandler) setupRoutes() {
	h.Router.HandleFunc("/books/", h.bookListHandler).Methods("GET")
	h.Router.HandleFunc("/books/{id}", h.bookGetHandler).Methods("GET")
	h.Router.HandleFunc("/books/", h.bookCreateHandler).Methods("POST")
	h.Router.HandleFunc("/books/{id}", h.bookUpdateHandler).Methods("PUT")
	h.Router.HandleFunc("/books/{id}", h.bookListHandler).Methods("DELETE")
}
