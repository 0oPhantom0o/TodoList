package DataBase

import (
	"TodoList/constants"
	"TodoList/domain"
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dsn := constants.Postgres
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database !")

	}
	fmt.Println("Database connected")
	err = DBConn.AutoMigrate(&domain.Todo{})
	if err != nil {
		return
	}
	fmt.Println("Migrated")
}
