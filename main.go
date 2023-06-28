package main

import (
	// "net/http"
	// "strconv"

	"log"
	"proyek/configs"
	"proyek/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// type Hello struct {
// 	Name string
// }

func main() {
	// loadEnv()
	configs.InitDB()
	e := echo.New()
	e = routes.InitRoute(e)
	//
	// e.GET("/hello", HelloController) //endpoint
	//route

	e.Start(":8000")
}

// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }
