package controllers

import (
	"github.com/camli1803/echo_golang/models"
)

// dau vao cua GetTodos GetATodo UpdateATodo
type TodoFilter struct {
	UserID int `json: "userID"`
	ID     int `json: "ID"`
}

// dau ra cua GetTodos
type TodoCollection struct {
	Todos []models.Todo
}

// dau vao cua CreateATodo
type CreateTodoInput struct {
	UserID  int    `json: "userID"`
	Content string `json: "content"`
}
type UpdateATodoInput struct {
	Content string `json: "userID"`
	IsDone  bool   `json: "isDone"`
}

func GetTodos(todoFilter TodoFilter) (TodoCollection, error) {
	var todos []models.Todo
	todos, err := models.FindAllTodos(todoFilter.UserID)
	var todoCollection TodoCollection
	todoCollection.Todos = todos
	return todoCollection, err
}

func CreateTodos(createTodoInput CreateTodoInput) (models.Todo, error) {
	var todo models.Todo
	todo.Content = createTodoInput.Content
	todo.UserID = createTodoInput.UserID
	todo.IsDone = false
	todoCreated, err := models.CreateTodos(todo)
	return todoCreated, err
}

func GetATodo(todoFilter TodoFilter) (models.Todo, error) {
	todo, err := models.FindATodo(todoFilter.UserID, todoFilter.ID)
	return todo, err
}

func UpdateATodo(todoFilter TodoFilter, updateATodo UpdateATodoInput) (models.Todo, error) {
	var todoNew models.Todo
	todoNew.Content = updateATodo.Content
	todoNew.IsDone = updateATodo.IsDone
	todo, err := models.UpdateATodo(todoFilter.UserID, todoFilter.ID, todoNew)
	return todo, err
}

func DeleteATodo(todoFilter TodoFilter) error {
	err := models.DeleteATodo(todoFilter.UserID, todoFilter.ID)
	return err
}
