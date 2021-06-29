package services

import (
	"github.com/kardeepak/http-mux-example/books/entities"
)

type BookService interface {
	List() []entities.Book
	Get(int) (entities.Book, error)
	Create(entities.Book) (entities.Book, error)
	Update(entities.Book) (entities.Book, error)
	Delete(int) entities.Book
}
