package handlers

import (
	_"github.com/nunows/todo-api/go/models"
    "github.com/gin-gonic/gin"
)

type TodoHandler struct {}

func (th *TodoHandler) GetTodos(c *gin.Context) {
    c.JSON(200, gin.H{"status": "GetTodos"})
}

func (th *TodoHandler) GetTodo(c *gin.Context) {
    c.JSON(200, gin.H{"status": "GetTodo"})
}

func (th *TodoHandler) PostTodo(c *gin.Context) {
	c.JSON(200,gin.H{"status": "PostTodo"})
}

func (th *TodoHandler) PutTodo(c *gin.Context) {
	c.JSON(200,gin.H{"status": "PutTodo"})
}

func (th *TodoHandler) DeleteTodo(c *gin.Context) {
	c.JSON(200,gin.H{"status": "DeleteTodo"})
}
