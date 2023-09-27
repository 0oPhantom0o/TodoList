package models

import (
	"ToDoList/database"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "check your input", "data": err})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not create todo", "data": err})
	}
	return c.JSON(&todo)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Coud not find todo", "data": err})
	}
	return c.JSON(&todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	type UpdateTodo struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Coud not update todo", "data": err})
	}
	var updatedTodo UpdateTodo
	err = c.BodyParser(&updatedTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "reveiw input", "data": err})
	}
	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed
	db.Save(&todo)
	return c.JSON(&todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Can not find Todo", "data": err})
	}
	db.Delete(&todo)
	return c.SendStatus(200)
}
