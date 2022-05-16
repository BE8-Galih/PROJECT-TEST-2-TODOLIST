package utils

import (
	"fmt"
	"todolist/config"
	"todolist/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := config.AddConfig()

	rds := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(rds), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}
	return db
}
func Migrate() {
	db := InitDB()
	db.AutoMigrate(&entities.User{}, &entities.Todo{}, &entities.Project{})
}
