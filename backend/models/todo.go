package models

import "github.com/jinzhu/gorm"

type TodoModel struct {
	gorm.Model
	
	Title     string `gorm:"not null" json:"title"`
	Completed int    `json:"completed"`
}

// TableName
func (todo TodoModel) TableName() string {
	return "todo"
}

// basic crud
func (todo TodoModel) Create() error {
	return db.Create(&todo).Error
}

func (todo TodoModel) Delete() error {
	return db.Delete(&todo).Error
}

func (todo TodoModel) Update() error {
	return db.Save(&todo).Error
}

func (todo TodoModel) Get() (TodoModel, error) {
	return todo, db.First(&todo, todo.ID).Error
}

func (todo TodoModel) GetAll() (uint64, []TodoModel, error) {
	var (
		count    uint64
		todoList []TodoModel
		err      error
	)
	
	if err = db.Table(todo.TableName()).Count(&count).Error; err != nil {
		return count, todoList, err
	}
	
	if err = db.Find(&todoList).Error; err != nil {
		return count, todoList, err
	}
	
	return count, todoList, nil
}
