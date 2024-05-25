package models

import (
	"TodoList/DataBase"
	"TodoList/domain"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	db := DataBase.DBConn
	var todos []domain.Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := DataBase.DBConn
	var todo domain.Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "couldn't find todo ", "data": err})
	}
	return c.JSON(&todo)
}

func CreateTodo(c *fiber.Ctx) error {
	db := DataBase.DBConn
	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "check your input ", "data": err})
	}
	err := db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "check your input ", "data": err})
	}
	return c.JSON(&todo)
}

func UpdateTodo(c *fiber.Ctx) error {

	id := c.Params("id")
	db := DataBase.DBConn
	var todo domain.Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "couldn't find todo ", "data": err})
	}
	var updateTodo domain.UpdatedTodo
	err = c.BodyParser(&updateTodo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "review input ", "data": err})
	}

	todo.Title = updateTodo.Title
	todo.Complete = updateTodo.Complete

	db.Save(&todo)
	return c.JSON(&updateTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := DataBase.DBConn
	var todo domain.Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "couldn't delete todo ", "data": err})
	}
	db.Delete(&todo)
	return c.JSON(200)
}
