package services

import (
	"errors"

	"appdynamics.com/golang/http-mux/books/entities"
	"appdynamics.com/golang/http-mux/books/repository"
	"github.com/go-playground/validator"
)

type bookServiceImpl struct {
	repository repository.BookRepository
	validate   *validator.Validate
}

func NewBookService(repository repository.BookRepository) BookService {
	return bookServiceImpl{repository, validator.New()}
}

func (svc bookServiceImpl) List() []entities.Book {
	return svc.repository.List()
}
func (svc bookServiceImpl) Get(ID int) (entities.Book, error) {
	return svc.repository.Get(ID)
}

func (svc bookServiceImpl) Create(book entities.Book) (entities.Book, error) {
	if book.ID != 0 {
		return entities.Book{}, errors.New("New Book must ID as zero. ")
	}
	if err := svc.validate.Struct(book); err != nil {
		return entities.Book{}, err
	}
	return svc.repository.Create(book), nil
}

func (svc bookServiceImpl) Update(book entities.Book) (entities.Book, error) {
	return svc.repository.Update(book)
}

func (svc bookServiceImpl) Delete(bookID int) entities.Book {
	return svc.repository.Delete(bookID)
}
