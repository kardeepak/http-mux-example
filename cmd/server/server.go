package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/kardeepak/http-mux-example/books/handlers"
	"github.com/kardeepak/http-mux-example/books/repository"
	"github.com/kardeepak/http-mux-example/books/services"
)

func main() {
	bookRepository := repository.NewInMemoryRepository()
	bookService := services.NewBookService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	port := "9000"

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(bookHandler)

	log.Info("Server is running on port: ", port)

	http.ListenAndServe(":"+port, n)
}
