package controllers

import (
	"log"
	"net/http"
	config "proyek/configs"
	"proyek/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewsController(c echo.Context) error {
	// var negara = c.QueryParam("negara")
	// var sort  = c.QueryParam("sort")

	//untuk query database

	var data []models.News

	result := config.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	// data = append(data, News{1, "A", "B"})
	// data = append(data, News{2, "B", "C"})
	// data = append(data, News{1, negara, "B"})
	// data = append(data, News{2, sort, "C"})

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    data,
	})
}

func DetailNewsController(c echo.Context) error {
	config.InitDB()
	var id, _ = strconv.Atoi(c.Param("id"))
	// id := c.Param("id")
	// var data = News{id, "a", "b"}
	// var data []News
	var news models.News
	// result := DB.Find(&news,id).First(&news)
	// result := DB.First(&data, id)
	result := config.DB.Find(&news, "id = ?", id)
	// result := DB.Where("id = ?", c.Param("id")).First(&news)
	// result := DB.Where("id = ?", id).First(&news)
	// result := DB.Where("id = ?", id).First(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    news,
	})
}

func AddNewsController(c echo.Context) error {
	// title := c.FormValue("title")
	// content := c.FormValue("content")
	var news models.News //keybinding json
	c.Bind(&news)

	//
	result := config.DB.Create(&news)
	//jika data kosong
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	// var data = News{1, title, content}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		// Data: data,
		Data: news, //key binding
	})
}

func UpdateNewsController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")
	content := c.FormValue("content")
	var news models.News
	// var data = models.News{id, title, content}
	// result := config.DB.Where().Update(&data)
	result := config.DB.Find(&news, "id = ?", id)

	news.Title = title
	news.Content = content
	config.DB.Save(&news)
	// result := config.DB.Where("id = ?", id).Updates(models.News{Title: title, Content: content})
	// result := config.DB.Where("id = ?", id).Updates(&news)
	log.Println("result : ", news)
	// config.DB.Save(result)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}
	// result := config.DB.Save(&news{title: title, content: content})

	// var data = News{1, title, content}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		// Data: data,
		Data: news, //key binding
	})
}

func DeleteNewsController(c echo.Context) error{
	var id, _ = strconv.Atoi(c.Param("id"))
	var news models.News

	result := config.DB.Find(&news, "id = ?", id)
	config.DB.Delete(&news)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		// Data: data,
		Data: news, //key binding
	})
}

// func HelloController(c echo.Context) error{
// 	var hello = Hello{"Alterra"}
// 	return c.JSON(http.StatusOK, hello)
// }
