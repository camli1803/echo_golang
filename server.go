package main

import (
	"errors"
	"net/http"

	"github.com/camli1803/echo_golang/db"
	"github.com/camli1803/echo_golang/models"
	"github.com/camli1803/echo_golang/views"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json: "name"`
	Todos []models.Todo
}

type CreateTodoInput struct {
	//{"content":"study IT"}
	Content string `json: "content"`
}

type UpdateTodo struct {
	Content string `json: "content"`
	IsDone  bool   `json: "isDone"`
}

func main() {
	db, err := db.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	//tao bang Todo, User
	db.AutoMigrate(&models.Todo{}, &User{})

	e := echo.New()

	// CRUD todo
	e.GET("/users/:userID/todos", views.GetTodos)
	// tao todo moi
	e.POST("/users/:userID/todos", views.CreateTodos)

	e.GET("/users/:userID/todos/:id", views.GetATodo)

	e.PATCH("/users/:userID/todos/:id", views.UpdateATodo)

	e.DELETE("/users/:userID/todos/:id", views.DeleteATodo)

	e.GET("/users/:userID", func(c echo.Context) error {
		var user User
		userID := c.Param("userID")
		result := db.First(&user, userID)
		isNotFoundError := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if result.Error != nil {
			if isNotFoundError {
				return c.JSON(http.StatusNotFound, "User ID not exist!")
			} else {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}
		var todos []models.Todo
		resultFindTodos := db.Where("user_id = ?", userID).Find(&todos)
		if resultFindTodos.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		user.Todos = todos
		return c.JSON(http.StatusOK, user)
	})

	e.PATCH("/users/:userID", func(c echo.Context) error {
		var user User
		id := c.Param("userID")
		result := db.First(&user, id)
		isNotFoundError := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if result.Error != nil {
			if isNotFoundError {
				return c.JSON(http.StatusNotFound, "User ID not exist!")
			} else {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		result_update := db.Model(&user).Updates(&user)
		if result_update.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, user)
	})

	e.DELETE("/users/:userID", func(c echo.Context) error {
		var user User
		userID := c.Param("userID")
		result := db.Delete(&user, userID)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		var todo models.Todo
		resultDeleteTodo := db.Where("user_id = ?", userID).Delete(&todo)
		if resultDeleteTodo.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "Delete Successful!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
