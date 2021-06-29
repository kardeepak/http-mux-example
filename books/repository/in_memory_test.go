package repository

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kardeepak/http-mux-example/books/entities"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	count := 10

	repo := NewInMemoryRepository()
	for i := 1; i <= count; i++ {
		repo.Create(entities.Book{})
	}

	books := repo.List()
	bookIDs := make([]int, 0)
	expectedIDs := make([]int, 0)
	for idx, book := range books {
		bookIDs = append(bookIDs, book.ID)
		expectedIDs = append(expectedIDs, idx+1)
	}
	assert.ElementsMatch(t, expectedIDs, bookIDs)
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	count := 10

	repo := NewInMemoryRepository()
	for i := 1; i <= count; i++ {
		repo.Create(entities.Book{})
	}

	for i := 1; i <= count; i++ {
		book, err := repo.Get(i)
		assert.Equal(t, i, book.ID)
		assert.Nil(t, err)
	}

	_, err := repo.Get(count + 1)
	assert.NotNil(t, err)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	count := 10

	repo := NewInMemoryRepository()
	for i := 1; i <= count; i++ {
		book := repo.Create(entities.Book{})
		assert.Equal(t, i, book.ID)
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewInMemoryRepository()

	book := repo.Create(entities.Book{Title: "Title"})
	updatedBook := book
	updatedBook.Title = "UpdatedTitle"

	acutal, err := repo.Update(updatedBook)
	assert.Equal(t, updatedBook, acutal)
	assert.Nil(t, err)

	acutal, err = repo.Get(book.ID)
	assert.Equal(t, updatedBook, acutal)
	assert.Nil(t, err)

}

func TestUpdateNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewInMemoryRepository()

	book := entities.Book{ID: 1, Title: "Title"}

	_, err := repo.Update(book)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	count := 10

	repo := NewInMemoryRepository()
	for i := 1; i <= count; i++ {
		repo.Create(entities.Book{})
	}

	for i := 1; i <= count; i++ {
		book := repo.Delete(i)
		assert.Equal(t, i, book.ID)
	}

	for i := 1; i <= count; i++ {
		_, err := repo.Get(i)
		assert.NotNil(t, err)
	}
}
