package book

import (
	"githihub.com/rajeshj3/go-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Book is a structure representing a single book
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks will return all books
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// GetBook will return a single book
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(404).Send("No book found")
		return
	}
	c.JSON(book)
}

// CreateBook will create a new book
func CreateBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

// DeleteBook will delete a book
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(404).Send("No book found")
		return
	}
	db.Delete(&book)
	c.JSON("Book successfully deleted!")
}
