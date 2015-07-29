package main

import (
	_"fmt"
	"github.com/nunows/todo-api/go/handlers"
    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()
    r.GET("/todo/", handlers.GetTodos)
    r.GET("/todo/:id", handlers.GetTodo)
	r.POST("/todo/", handlers.PostTodo)
	r.PUT("/todo/:id", handlers.PutTodo)
	r.DELETE("/todo/:id", handlers.DeleteTodo)
    r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
