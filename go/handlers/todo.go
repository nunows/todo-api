package handlers

import (
	_"github.com/nunows/todo-api/go/models"
    "github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
    c.JSON(200, gin.H{"status": "GetTodos"})
}

func GetTodo(c *gin.Context) {
    c.JSON(200, gin.H{"status": "GetTodo"})
}

func PostTodo(c *gin.Context) {
	c.JSON(200,gin.H{"status": "PostTodo"})
}

func PutTodo(c *gin.Context) {
	c.JSON(200,gin.H{"status": "PutTodo"})
}

func DeleteTodo(c *gin.Context) {
	c.JSON(200,gin.H{"status": "DeleteTodo"})
}
