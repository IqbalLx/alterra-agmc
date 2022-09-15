package libs

import "static-crud/models"

func CreateBook(book models.Book) models.Book {
	return models.Book{
		Id:     1,
		Title:  "Atomic Habits",
		Author: "James Clear",
	}
}

func GetBook(id int) models.Book {
	return models.Book{
		Id:     id,
		Title:  "Atomic Habits",
		Author: "James Clear",
	}
}

func GetBooks() []models.Book {
	return []models.Book{
		{
			Id:     1,
			Title:  "Atomic Habits",
			Author: "James Clear",
		},
		{
			Id:     2,
			Title:  "Think and Grow Rich",
			Author: "Napolean Hills",
		},
		{
			Id:     3,
			Title:  "Clean Architecture",
			Author: "Robert C. Martin",
		},
	}
}

func UpdateBook(id int, book models.Book) models.Book {
	return book
}

func DeleteBook(id int) models.Book {
	return models.Book{
		Id:     id,
		Title:  "Atomic Habits",
		Author: "James Clear",
	}
}
