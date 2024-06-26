package main

import (
	"log"
	"os"
	"rest-api-postgres/database"
	"rest-api-postgres/handlers"
	"rest-api-postgres/routes"
	"rest-api-postgres/services/book"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	defer database.CloseDB()

	bookService := book.NewBookService(database.DB)
	bookHandler := handlers.NewBookHandler(bookService)

	// Set gin mode base on environment variable
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = gin.ReleaseMode // Default to release mode if not exist
	}

	gin.SetMode(mode)
	// Initialize the router
	r := gin.Default()
	// Set trusted proxies
	trustedProxies := os.Getenv("TRUSTED_PROXIES")
	if trustedProxies != "" {
		r.SetTrustedProxies(strings.Split(trustedProxies, ","))
	}
	// Define your routes
	api := r.Group("api/")
	{
		routes.BookRouter(api, bookHandler)
	}

	listenPort := ":" + os.Getenv("APP_PORT")
	r.Run(listenPort)
}
