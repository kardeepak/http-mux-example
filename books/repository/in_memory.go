package repository

import (
	"errors"
	"fmt"

	"github.com/kardeepak/http-mux-example/books/entities"
)

type inMemoryBookRepository struct {
	books map[int]entities.Book
}

func NewInMemoryRepository() BookRepository {
	return inMemoryBookRepository{
		books: make(map[int]entities.Book),
	}
}

func (repo inMemoryBookRepository) List() []entities.Book {
	books := make([]entities.Book, 0)
	for _, book := range repo.books {
		books = append(books, book)
	}
	return books
}

func (repo inMemoryBookRepository) Get(ID int) (entities.Book, error) {
	if _, ok := repo.books[ID]; !ok {
		msg := fmt.Sprintf("Book with ID : %d doesn't exist.", ID)
		return entities.Book{}, errors.New(msg)
	}
	return repo.books[ID], nil
}

func (repo inMemoryBookRepository) Create(book entities.Book) entities.Book {
	book.ID = len(repo.books) + 1
	repo.books[book.ID] = book
	return book
}

func (repo inMemoryBookRepository) Update(book entities.Book) (entities.Book, error) {
	if _, ok := repo.books[book.ID]; !ok {
		msg := fmt.Sprintf("Book with ID : %d doesn't exist.", book.ID)
		return entities.Book{}, errors.New(msg)
	}
	repo.books[book.ID] = book
	return book, nil
}

func (repo inMemoryBookRepository) Delete(bookID int) entities.Book {
	book := repo.books[bookID]
	delete(repo.books, bookID)
	return book
}
