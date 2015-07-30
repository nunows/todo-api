package main

import (
	_"fmt"
	"github.com/nunows/todo-api/go/handlers"
    "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	_"github.com/nunows/todo-api/go/models"
)



func main() {

	var db gorm.DB

	db, _ = gorm.Open("sqlite3", "../database/todo.db")
	db.LogMode(true)
	//db.AutoMigrate(&models.Todo{})

	th := handlers.TodoHandler{Db: db}

    r := gin.Default()
    r.GET("/todo/", th.GetAll)
    r.GET("/todo/:id", th.Get)
	r.POST("/todo/", th.Create)
	r.PUT("/todo/:id", th.Update)
	r.DELETE("/todo/:id", th.Delete)
    r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
