package handlers

import (
	"fmt"
	"net/http"
	"rest-api-postgres/helpers"
	"rest-api-postgres/models"
	"rest-api-postgres/services/book"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service book.Service
}

func NewBookHandler(s book.Service) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

func (h *BookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.service.FetchBooks()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   true,
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"error":   false,
		"message": "success",
		"data":    books,
	})
}

func (h *BookHandler) AddBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   true,
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		})
		return
	}

	_, err := h.service.InserBook(newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   true,
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"error":   false,
		"message": "Successfully added data!",
		"data":    nil,
	})
}

func (h *BookHandler) GetBook(ctx *gin.Context) {
	Id, err := helpers.Str2Int(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   true,
			"message": "Book with ID invalid syntax!",
			"data":    nil,
		})
		return
	}

	book, err := h.service.FetchBook(Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   true,
			"message": fmt.Sprintf("Book with ID %d not found", Id),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"error":   false,
		"message": "success",
		"data":    book,
	})
}

func (h *BookHandler) UpdateBook(ctx *gin.Context) {
	var book models.Book

	Id, err := helpers.Str2Int(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   true,
			"message": "Book with ID invalid syntax!",
			"data":    nil,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   true,
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		})
		return
	}

	_, err = h.service.FetchBook(Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   true,
			"message": fmt.Sprintf("Book with ID %d not found", Id),
			"data":    nil,
		})
		return
	}

	_, err = h.service.UpdateBook(Id, book)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   true,
			"message": "Data failed updated!",
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"error":   false,
		"message": "Data successfully updated!",
		"data":    nil,
	})
}

func (h *BookHandler) RemoveBook(ctx *gin.Context) {
	Id, err := helpers.Str2Int(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   true,
			"message": "Book with ID invalid syntax!",
			"data":    nil,
		})
		return
	}

	_, err = h.service.FetchBook(Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   true,
			"message": fmt.Sprintf("Book with ID %d not found", Id),
			"data":    nil,
		})
		return
	}

	err = h.service.RemoveBook(Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   true,
			"message": "Data failed deleted!",
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"error":   false,
		"message": "Data successfully deleted!",
		"data":    nil,
	})
}
