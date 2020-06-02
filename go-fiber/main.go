package main

import (
	"fmt"

	"githihub.com/rajeshj3/go-fiber/book"
	"githihub.com/rajeshj3/go-fiber/database"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/", book.CreateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Fail to connect to database!")
	}
	fmt.Println("Database connected!")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated!")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen("localhost:3000")
}
