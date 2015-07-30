package handlers

import (
	"github.com/nunows/todo-api/go/models"
    "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type TodoHandler struct {
	Db gorm.DB
}

func (th *TodoHandler) GetAll(c *gin.Context) {
	var todos models.Todos
    th.Db.Find(&todos)
    c.JSON(200, todos)
}

func (th *TodoHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if (th.Db.First(&todo, id).RecordNotFound()) {
		c.JSON(404, gin.H{"error:": "Todo not found"})
	} else {
		c.JSON(200, todo)
	}
}

func (th *TodoHandler) Create(c *gin.Context) {
	var todo models.Todo

	if (c.BindJSON(&todo) == nil) {
		th.Db.Create(&todo)
		c.JSON(200,gin.H{"status": "Todo created"})
	} else {
		c.JSON(500, gin.H{"error:": "Creating todo"})
	}
}

func (th *TodoHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var todoDb models.Todo
	var todo models.Todo

	if (th.Db.First(&todoDb, id).RecordNotFound()) {
		c.JSON(404, gin.H{"error:": "Todo not found"})
	} else {
		if (c.BindJSON(&todo) == nil) {
			todoDb.Name = todo.Name
			todoDb.Done = todo.Done
			th.Db.Save(&todoDb)
			c.JSON(200,gin.H{"status": "Todo updated"})
		} else {
			c.JSON(500, gin.H{"error:": "Updating todo"})
		}
	}
}

func (th *TodoHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if (th.Db.First(&todo, id).RecordNotFound()) {
		c.JSON(404, gin.H{"error:": "Todo not found"})
	} else {
		th.Db.Delete(&todo)
		c.JSON(200,gin.H{"status": "Todo deleted"})
	}
}
