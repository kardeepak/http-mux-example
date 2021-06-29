package repository

import "github.com/kardeepak/http-mux-example/books/entities"

//go:generate mockgen -source $GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
type BookRepository interface {
	List() []entities.Book
	Get(int) (entities.Book, error)
	Create(entities.Book) entities.Book
	Update(entities.Book) (entities.Book, error)
	Delete(int) entities.Book
}
