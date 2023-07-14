package database

import (
	"fmt"
	"go-curd-redis-example/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB(config *config.Config) *gorm.DB {
	dsn1 := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	fmt.Println(dsn1)
	fmt.Println(config.DBUsername)
	dsn := "root:@tcp(127.0.0.1:3306)/go_fiber_practice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	fmt.Println("Connect db success!!!!")
	return db
}
