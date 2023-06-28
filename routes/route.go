package routes

import (
	"proyek/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {

	e.GET("/news", controllers.NewsController)           //endpoint
	e.GET("/news/:id", controllers.DetailNewsController) //endpoint
	e.POST("/news", controllers.AddNewsController)       //endpoint
	e.PUT("/news/:id", controllers.UpdateNewsController)       //endpoint
	e.DELETE("/news/:id", controllers.DeleteNewsController)       //endpoint
	return e
}
