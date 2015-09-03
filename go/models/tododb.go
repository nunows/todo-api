package models

import (
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
	td.Db.LogMode(false)
	td.Db.AutoMigrate(&Todo{})
	return nil
}

func (td *TodoDb) GetAll() (*Todos, error) {
	var todos Todos
	err := td.Db.Find(&todos).Error
	if err != nil && err != gorm.RecordNotFound {
		return nil, err
	}
	return &todos, nil
}

func (td *TodoDb) Get(id int) (*Todo, error) {
	var todo Todo

	if td.Db.First(&todo, id).RecordNotFound() {
		return nil, gorm.RecordNotFound
	}
	return &todo, nil
}

func (td *TodoDb) Insert(todo *Todo) error {
	err := td.Db.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (td *TodoDb) Update(id int, todo *Todo) error {
	update, err := td.Get(id)
	if err != nil {
		return err
	}

	update.Name = todo.Name
	update.Done = todo.Done
	err = td.Db.Save(&update).Error
	if err != nil {
		return err
	}
	return nil

}

func (td *TodoDb) Delete(id int) error {
	delete, err := td.Get(id)
	if err != nil {
		return err
	}

	err = td.Db.Delete(&delete).Error
	if err != nil {
		return err
	}
	return nil
}
