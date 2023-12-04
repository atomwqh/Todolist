package Models

import (
	"bubble_todolist/DAO"
)

// todolist项目model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查

func CreateATodo(todo *Todo) (err error) {
	err = DAO.Db.Create(&todo).Error
	return
}

func GetALLTodo() (todolist []*Todo, err error) {
	if err := DAO.Db.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = DAO.Db.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}
func UpdateATodo(todo *Todo) (err error) {
	err = DAO.Db.Save(todo).Error
	return
}
func DeleteATodo(id string) (err error) {
	err = DAO.Db.Where("id=?", id).Delete(&Todo{}).Error
	return
}
