package models

import (
	"fmt"

	"github.com/camli1803/echo_golang/db"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `json: "content"`
	IsDone  bool   `json: "isDone"`
	UserID  int    `json: "userID"`
}

// TODO: co bug
func FindAllTodos(userID int) ([]Todo, error) {
	var todos []Todo
	result := db.DB.Where("user_id = ?", userID).Find(&todos)
	if result.Error != nil {
		err := fmt.Errorf("Internal Server Error")
		return nil, err
	}
	return todos, nil
}

func CreateTodos(todo Todo) (Todo, error) {
	result := db.DB.Create(&todo)
	if result.Error != nil {
		err := fmt.Errorf("Internet Server Error")
		// return todo rong nhung chua tim duoc
		return todo, err
	}
	return todo, nil
}

func FindATodo(userID, id int) (todo Todo, err error) {
	var todos []Todo
	result := db.DB.Where("user_id = ? AND id = ?", userID, id).Find(&todos)
	if result.Error != nil {
		err := fmt.Errorf("Internet Server Error")
		// chua tim duoc gia tri rong cho todo
		return todo, err
	}
	if len(todos) == 0 {
		err := fmt.Errorf("UserID or id not exist")
		return todo, err
	}
	return todos[0], nil
}

func UpdateATodo(userID, id int, todoNew Todo) (Todo, error) {
	var todo Todo
	todo, err := FindATodo(userID, id)
	if err != nil {
		err := fmt.Errorf("Error!")
		return todo, err
	}
	result := db.DB.Model(&todo).Updates(&todoNew)
	if result.Error != nil {
		err := fmt.Errorf("Internal Server Error")
		return todo, err
	}
	return todo, nil
}

func DeleteATodo(userID, id int) error {
	var todos []Todo
	result := db.DB.Where("user_id = ? AND id = ?", userID, id).Find(&todos)
	if result.Error != nil {
		err := fmt.Errorf("Internet Server Error")
		return err
	}
	if len(todos) == 0 {
		err := fmt.Errorf("UserID or id not exist")
		return err
	}
	resultDelete := db.DB.Delete(&todos[0], id)
	if resultDelete.Error != nil {
		err := fmt.Errorf("Internet Server Error")
		return err
	}
	return nil
}
