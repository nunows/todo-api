package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nunows/todo-api/go/handlers"
	"github.com/nunows/todo-api/go/models"
)

func main() {

	db := new(models.TodoDb)
	db.Open()

	th := handlers.TodoHandler{Db: *db}

	r := gin.Default()
	r.GET("/todo/", th.GetAll)
	r.GET("/todo/:id", th.Get)
	r.POST("/todo/", th.Create)
	r.PUT("/todo/:id", th.Update)
	r.DELETE("/todo/:id", th.Delete)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
