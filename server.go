package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `json: "content"`
	IsDone  bool   `json: "isDone"`
}

func main() {

	db, err := gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{})

	e := echo.New()

	e.GET("/todos", func(c echo.Context) error {
		// tra ve 1 mang todo
		var todos []Todo
		result := db.Find(&todos)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, todos)
	})

	// tao todo moi
	e.POST("/todos", func(c echo.Context) error {
		//	lay du lieu tu request
		//	content := c.FormValue("content")
		todo := Todo{Content: "", IsDone: false}
		// ham Bind lay du lieu tu request
		if err := c.Bind(&todo); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		result := db.Create(&todo)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusCreated, todo)
	})

	e.GET("/todos/:id", func(c echo.Context) error {
		var todo Todo
		id := c.Param("id")
		result := db.First(&todo, id)
		isNotFoundError := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if result.Error != nil {
			if isNotFoundError {
				return c.JSON(http.StatusNotFound, "ID not exist!")
			} else {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}
		return c.JSON(http.StatusOK, todo)
	})

	e.PATCH("/todos/:id", func(c echo.Context) error {
		var todo Todo
		id := c.Param("id")
		result := db.First(&todo, id)
		isNotFoundError := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if result.Error != nil {
			if isNotFoundError {
				return c.JSON(http.StatusNotFound, "ID not exist!")
			} else {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}
		if err := c.Bind(&todo); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		result_update := db.Model(&todo).Updates(&todo)
		if result_update.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, todo)
	})

	e.DELETE("todos/:id", func(c echo.Context) error {
		var todo Todo
		id := c.Param("id")
		result := db.Delete(&todo, id)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "Delete Successful!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
