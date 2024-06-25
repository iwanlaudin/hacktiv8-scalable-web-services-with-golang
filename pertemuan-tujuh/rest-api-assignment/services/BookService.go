package services

import "rest-api-assignment/models"

func FindAll() []models.Book {
	return models.BookDatas
}

func Find(Id int) (models.Book, int, bool) {
	var book models.Book
	var condition = false
	var index int

	for i, b := range models.BookDatas {
		if Id == b.Id {
			book = models.BookDatas[i]
			condition = true
			index = i
			break
		}
	}

	return book, index, condition
}

func Add(collection models.Book) models.Book {
	collection.Id = len(models.BookDatas) + 1
	models.BookDatas = append(models.BookDatas, collection)

	return collection
}

func Updated(id int, index int, collection models.Book) models.Book {
	collection.Id = id
	models.BookDatas[index] = collection
	return collection
}

func Deleted(index int) {
	copy(models.BookDatas[index:], models.BookDatas[index+1:])
	models.BookDatas[len(models.BookDatas)-1] = models.Book{}
	models.BookDatas = models.BookDatas[:len(models.BookDatas)-1]
}
