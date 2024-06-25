package routes

import (
	"rest-api-postgres/handlers"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.RouterGroup, h *handlers.BookHandler) {
	router.GET("/books", h.GetBooks)
	router.POST("/books", h.AddBook)
	router.GET("/books/:id", h.GetBook)
	router.PUT("/books/:id", h.UpdateBook)
	router.DELETE("/books/:id", h.RemoveBook)
}
