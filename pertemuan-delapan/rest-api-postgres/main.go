package main

import (
	"rest-api-postgres/config"
	"rest-api-postgres/database"
	"rest-api-postgres/handlers"
	"rest-api-postgres/routes"
	"rest-api-postgres/services/book"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	bookService := book.NewBookService(database.DB)
	bookHandler := handlers.NewBookHandler(bookService)

	r := gin.Default()
	api := r.Group("api/")
	{
		routes.BookRouter(api, bookHandler)
	}

	r.Run(config.PORT)
}
