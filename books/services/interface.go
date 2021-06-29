package services

import (
	"appdynamics.com/golang/http-mux/books/entities"
)

type BookService interface {
	List() []entities.Book
	Get(int) (entities.Book, error)
	Create(entities.Book) (entities.Book, error)
	Update(entities.Book) (entities.Book, error)
	Delete(int) entities.Book
}
