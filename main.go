package main

import (
	"TodoList/DataBase"
	"TodoList/constants"
	"TodoList/models"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get(constants.All, models.GetTodos)
	app.Post(constants.Create, models.CreateTodo)
	app.Get(constants.GetOne, models.GetTodoById)
	app.Put(constants.Change, models.UpdateTodo)
	app.Delete(constants.Delete, models.DeleteTodo)
}
func main() {

	DataBase.InitDatabase()
	app := fiber.New()
	setupRoutes(app)
	err := app.Listen(constants.Port)
	if err != nil {
		return
	}
}
