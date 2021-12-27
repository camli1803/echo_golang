package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `json: "content"`
	IsDone  bool   `json: "isDone"`
	UserID  int    `json: "userID"`
}

type User struct {
	gorm.Model
	Name  string `json: "name"`
	Todos []Todo
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

	db, err := gorm.Open(sqlite.Open("todolist1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//tao bang Todo, User
	db.AutoMigrate(&Todo{}, &User{})

	e := echo.New()

	// CRUD todo
	e.GET("/users/:userID/todos", func(c echo.Context) error {
		// tra ve 1 mang todo
		var todos []Todo
		userID := c.Param("userID")
		result := db.Where("user_id = ?", userID).Find(&todos)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		if len(todos) == 0 {
			return c.JSON(http.StatusNotFound, "User ID not exist")
		}
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, todos)
	})

	// tao todo moi
	e.POST("/users/:userID/todos", func(c echo.Context) error {
		//lay userID
		var user User
		userID := c.Param("userID")
		//kiem tra user ID co ton tai khong
		resultFindUserID := db.First(&user, userID)
		isNotFoundError := errors.Is(resultFindUserID.Error, gorm.ErrRecordNotFound)
		if resultFindUserID.Error != nil {
			if isNotFoundError {
				return c.JSON(http.StatusNotFound, "User ID not exist!")
			} else {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}
		iUserID, _ := strconv.Atoi(userID)
		//	lay du lieu tu request
		todoInput := CreateTodoInput{Content: ""}
		// ham Bind lay du lieu tu request
		if err := c.Bind(&todoInput); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		todo := Todo{Content: todoInput.Content, IsDone: false, UserID: iUserID}
		result := db.Create(&todo)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusCreated, todo)
	})

	e.GET("/users/:userID/todos/:id", func(c echo.Context) error {
		userID := c.Param("userID")
		var todos []Todo
		id := c.Param("id")
		result := db.Where("user_id = ? AND id = ?", userID, id).Find(&todos)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		if len(todos) == 0 {
			return c.JSON(http.StatusNotFound, "User ID or todo id not exist!")
		}
		return c.JSON(http.StatusOK, todos[0])
	})

	e.PATCH("/users/:userID/todos/:id", func(c echo.Context) error {
		userID := c.Param("userID")
		iUserID, _ := strconv.Atoi(userID)
		var todos []Todo
		id := c.Param("id")
		result := db.Where("user_id = ? AND id = ?", userID, id).Find(&todos)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		if len(todos) == 0 {
			return c.JSON(http.StatusNotFound, "User ID or todo id not exist!")
		}
		var updatetodo UpdateTodo
		if err := c.Bind(&updatetodo); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		todo := Todo{Content: updatetodo.Content, IsDone: updatetodo.IsDone, UserID: iUserID}
		result_update := db.Model(&todos[0]).Updates(&todo)
		if result_update.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, todos[0])
	})

	e.DELETE("/users/:userID/todos/:id", func(c echo.Context) error {
		userID := c.Param("userID")
		var todos []Todo
		id := c.Param("id")
		result := db.Where("user_id = ? AND id = ?", userID, id).Find(&todos)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		if len(todos) == 0 {
			return c.JSON(http.StatusNotFound, "User ID or todo id not exist!")
		}
		resultDelete := db.Delete(&todos[0], id)
		if resultDelete.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "Delete Successful!")
	})

	// CRUD user
	e.GET("/users", func(c echo.Context) error {
		var users []User
		result := db.Find(&users)
		for i := 0; i < len(users); i++ {
			var todos []Todo
			resultFindTodos := db.Where("user_id = ?", users[i].ID).Find(&todos)
			if resultFindTodos.Error != nil {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
			users[i].Todos = todos
		}
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		var user User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		result := db.Create(&user)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusCreated, user)

	})

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
		var todos []Todo
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
		var todo Todo
		resultDeleteTodo := db.Where("user_id = ?", userID).Delete(&todo)
		if resultDeleteTodo.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "Delete Successful!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
