package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type TodoDb struct {
	Db gorm.DB
}

func (td *TodoDb) Open(dbpath string) error {
	var err error
	td.Db, err = gorm.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}
	td.Db.LogMode(true)
	td.Db.AutoMigrate(&Todo{})
	return nil
}

func (td *TodoDb) GetAll() Todos {
	var todos Todos
	td.Db.Find(&todos)
	return todos
}

func (td *TodoDb) Get(id int) (Todo, error) {
	var todo Todo

	if td.Db.First(&todo, id).RecordNotFound() {
		return todo, errors.New("Record Not Found")
	} else {
		return todo, nil
	}
}

func (td *TodoDb) Insert(todo *Todo) bool {
	td.Db.Create(&todo)
	return true
}

func (td *TodoDb) Update(id int, todo *Todo) bool {
	update, err := td.Get(id)

	if err == nil {
		update.Name = todo.Name
		update.Done = todo.Done
		td.Db.Save(&update)
		return true
	} else {
		return false
	}
}

func (td *TodoDb) Delete(id int) bool {
	delete, err := td.Get(id)

	if err == nil {
		td.Db.Delete(&delete)
		return true
	} else {
		return false
	}
}
