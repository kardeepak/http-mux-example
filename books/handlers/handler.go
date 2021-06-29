package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"appdynamics.com/golang/http-mux/books/entities"
	"appdynamics.com/golang/http-mux/books/services"
)

type BookHandler struct {
	service services.BookService
	*mux.Router
}

// GET /books/{id}
func (h BookHandler) bookGetHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("bookGetHandler")
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := h.service.Get(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// GET /books
func (h BookHandler) bookListHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("bookListHandler")
	json.NewEncoder(w).Encode(h.service.List())
}

// POST /books
func (h BookHandler) bookCreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("bookCreateHandler")
	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "test", http.StatusBadRequest)
		return
	}
	book, err := h.service.Create(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// PUT /books/{id}
func (h BookHandler) bookUpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("bookUpdateHandler")
	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.ID = bookID
	book, err = h.service.Update(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// DELETE /books/{id}
func (h BookHandler) bookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("bookDeleteHandler")
	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book := h.service.Delete(bookID)
	json.NewEncoder(w).Encode(book)
}
