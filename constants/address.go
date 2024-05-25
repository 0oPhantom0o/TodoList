package constants

const (
	All    = "/todos/all"
	Create = "/todos"
	GetOne = "/todos-get/:id"
	Change = "/todos-change/:id"
	Delete = "/todos-delete/:id"
)

const Port = ":8080"

const Postgres = "host=localhost port=5432 user=postgres dbname=phantom sslmode=disable password=postgres"
