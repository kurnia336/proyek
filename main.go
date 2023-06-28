package main

import (
	// "net/http"
	// "strconv"

	"os"
	"proyek/configs"
	"proyek/routes"

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

	// e.Start(":8000")
	e.Start(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }
