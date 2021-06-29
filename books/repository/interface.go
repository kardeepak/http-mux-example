package repository

import "appdynamics.com/golang/http-mux/books/entities"

type BookRepository interface {
	List() []entities.Book
	Get(int) (entities.Book, error)
	Create(entities.Book) entities.Book
	Update(entities.Book) (entities.Book, error)
	Delete(int) entities.Book
}
