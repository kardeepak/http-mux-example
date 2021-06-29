package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/kardeepak/http-mux-example/books/entities"
	"github.com/kardeepak/http-mux-example/books/repository"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	bookList := []entities.Book{{ID: 1, Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 9.99}}
	mockRepo.EXPECT().List().Return(bookList)

	svc := NewBookService(mockRepo)

	assert.Equal(t, bookList, svc.List())
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	book := entities.Book{ID: 1, Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 9.99}
	mockRepo.EXPECT().Get(1).Return(book, nil)

	svc := NewBookService(mockRepo)

	actual, err := svc.Get(1)
	assert.Equal(t, book, actual)
	assert.Nil(t, err)
}

func TestGetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	err := errors.New("Book with ID : 1 does't exist.")
	mockRepo.EXPECT().Get(1).Return(entities.Book{}, err)

	svc := NewBookService(mockRepo)

	_, err = svc.Get(1)
	assert.NotNil(t, err)
	t.Log("Error : ", err.Error())
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	book := entities.Book{Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 9.99}
	savedBook := entities.Book(book)
	savedBook.ID = 1
	mockRepo.EXPECT().Create(book).Return(savedBook)

	svc := NewBookService(mockRepo)

	actual, err := svc.Create(book)
	assert.Equal(t, savedBook, actual)
	assert.Nil(t, err)
}

func TestCreateNonZeroID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	book := entities.Book{ID: 1, Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 9.99}

	svc := NewBookService(mockRepo)

	_, err := svc.Create(book)
	assert.NotNil(t, err)
	t.Log("Error : ", err.Error())
}

func TestCreateInvalidPrice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	book := entities.Book{Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 2.99}

	svc := NewBookService(mockRepo)

	_, err := svc.Create(book)
	assert.NotNil(t, err)
	t.Log("Error : ", err.Error())
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	book := entities.Book{ID: 1, Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 9.99}

	updatedBook := book
	updatedBook.Price = 19.99
	mockRepo.EXPECT().Update(book).Return(updatedBook, nil)

	svc := NewBookService(mockRepo)

	actual, err := svc.Update(book)
	assert.Equal(t, updatedBook, actual)
	assert.Nil(t, err)
}

func TestUpdateNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	book := entities.Book{ID: 1, Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 9.99}

	err := errors.New("Book with ID : 1 does't exist.")
	mockRepo.EXPECT().Update(book).Return(entities.Book{}, err)

	svc := NewBookService(mockRepo)

	_, err = svc.Update(book)
	assert.NotNil(t, err)
	t.Log("Error : ", err.Error())
}

func TestUpdateInvalidPrice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockBookRepository(ctrl)

	book := entities.Book{ID: 1, Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 2.99}

	svc := NewBookService(mockRepo)

	_, err := svc.Update(book)
	assert.NotNil(t, err)
	t.Log("Error : ", err.Error())
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	deletedBook := entities.Book{ID: 1, Title: "Title", Author: "Author",
		ISBN: "ISBN", Description: "Description", Price: 2.99}

	mockRepo := repository.NewMockBookRepository(ctrl)
	mockRepo.EXPECT().Delete(1).Return(deletedBook)

	svc := NewBookService(mockRepo)

	actual := svc.Delete(1)
	assert.Equal(t, deletedBook, actual)
}
