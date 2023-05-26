package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Migration() {
	errs := godotenv.Load(".env")
	if errs != nil {
		log.Fatal("Error Loading Env File: ", err)
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v  dbname=%v  port=5432 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("not connected")
	} else {
		fmt.Println("connected")
	}
	DB.AutoMigrate(&User{})

}

func main() {
	Migration()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("hello world")
	})
	app.Post("/users", func(c *fiber.Ctx) error {
		var user struct {
			Name string `json:"name"`
			Age  string `json:"age"`
		}
		c.BodyParser(&user)
		fmt.Println(user)
		return c.JSON(&user)
	})
	app.Listen(":3000")
}

type User struct {
	Name     string
	Lastname string
}
