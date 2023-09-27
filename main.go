package main

import (
	"fmt"

	"ToDoList/database"
	"ToDoList/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func hellowWord(c *fiber.Ctx) error {
	return c.SendString("hello world!!!!!!!!!!!!!!!!!!!!")
}

func initDatabase() {
	var err error

	dsn := "host=localhost user=postgres password=12345678 dbname=postgres port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open database!")

	}
	fmt.Println("Database Connection established")
	database.DBConn.AutoMigrate((&models.Todo{}))
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.CreateTodo)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}
func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	app.Get("/", hellowWord)
	setupRoutes(app)
	app.Listen(":8000")
}
