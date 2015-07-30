package main

import (
	_"fmt"
	"github.com/nunows/todo-api/go/handlers"
    "github.com/gin-gonic/gin"
)

func main() {

	th := handlers.TodoHandler{}

    r := gin.Default()
    r.GET("/todo/", th.GetTodos)
    r.GET("/todo/:id", th.GetTodo)
	r.POST("/todo/", th.PostTodo)
	r.PUT("/todo/:id", th.PutTodo)
	r.DELETE("/todo/:id", th.DeleteTodo)
    r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
