package main

import (
	"fmt"
	"go-curd-redis-example/config"
	"go-curd-redis-example/database"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello world")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load evirontmetn variabe", err)
	}

	// mysql
	db := database.ConnectionDB(&loadConfig)
	rdb := database.ConnectionRedisDb(&loadConfig)

	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()

	err := app.Listen(":3400")
	if err != nil {
		panic(err)
	}
}
