package views

import (
	"net/http"
	"strconv"

	"github.com/camli1803/echo_golang/controllers"
	"github.com/camli1803/echo_golang/models"
	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	var todoFilter controllers.TodoFilter
	userID := c.Param("userID")
	iuserID, _ := strconv.Atoi(userID)
	todoFilter.UserID = iuserID
	var todoCollection controllers.TodoCollection
	todoCollection, err := controllers.GetTodos(todoFilter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, todoCollection)
}

func CreateTodos(c echo.Context) error {
	var createTodoInput controllers.CreateTodoInput
	var todo models.Todo
	userID := c.Param("userID")
	iuserID, _ := strconv.Atoi(userID)
	if err := c.Bind(&createTodoInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	createTodoInput.UserID = iuserID
	todo, err := controllers.CreateTodos(createTodoInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, todo)
}

func GetATodo(c echo.Context) error {
	// dau vao cho controller
	var todoFilter controllers.TodoFilter
	userID := c.Param("userID")
	iuserID, _ := strconv.Atoi(userID)
	id := c.Param("id")
	iID, _ := strconv.Atoi(id)
	todoFilter.UserID = iuserID
	todoFilter.ID = iID

	//dau ra cho controller
	var todo models.Todo
	todo, err := controllers.GetATodo(todoFilter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todo)
}

func UpdateATodo(c echo.Context) error {
	//dau vao cua controller
	var todoFilter controllers.TodoFilter
	userID := c.Param("userID")
	iuserID, _ := strconv.Atoi(userID)
	id := c.Param("id")
	iID, _ := strconv.Atoi(id)
	todoFilter.UserID = iuserID
	todoFilter.ID = iID

	var updateATodo controllers.UpdateATodoInput
	if err := c.Bind(&updateATodo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todo, err := controllers.UpdateATodo(todoFilter, updateATodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todo)
}

func DeleteATodo(c echo.Context) error {
	var todoFilter controllers.TodoFilter
	userID := c.Param("userID")
	iuserID, _ := strconv.Atoi(userID)
	id := c.Param("id")
	iID, _ := strconv.Atoi(id)
	todoFilter.UserID = iuserID
	todoFilter.ID = iID
	err := controllers.DeleteATodo(todoFilter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Delete Successful")
}
