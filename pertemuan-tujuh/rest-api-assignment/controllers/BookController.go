package controllers

import (
	"net/http"
	"rest-api-assignment/models"
	"rest-api-assignment/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	books := services.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": books,
	})
}

func Create(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook = services.Add(newBook)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    newBook,
		"message": "Successfully added data!",
	})
}

func Detail(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book, _, condition = services.Find(Id)

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": book,
	})
}

func Updated(ctx *gin.Context) {
	var updatedBook models.Book

	id, err := strconv.Atoi(ctx.Param("Id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		return
	}

	var _, index, condition = services.Find(id)

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	services.Updated(id, index, updatedBook)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Data successfully updated!",
	})
}

func Deleted(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("Id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		return
	}

	var _, index, condition = services.Find(id)

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data not found!",
		})
		return
	}

	services.Deleted(index)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Data successfully deleted!",
	})
}
