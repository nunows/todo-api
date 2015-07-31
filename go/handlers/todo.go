package handlers

import (
	"strconv"
	"github.com/nunows/todo-api/go/models"
    "github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Db models.TodoDb
}

func (th *TodoHandler) GetAll(c *gin.Context) {
    c.JSON(200, th.Db.GetAll())
}

func (th *TodoHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := th.Db.Get(id)

    if(err == nil){
		c.JSON(200, todo)
	} else {
		c.JSON(404, gin.H{"error:": "Todo not found"})
	}
}

func (th *TodoHandler) Create(c *gin.Context) {
	var todo models.Todo

	if (c.BindJSON(&todo) == nil) {
		//check bool
		th.Db.Insert(&todo)
		c.JSON(200,gin.H{"status": "Todo created"})
	} else {
		c.JSON(500, gin.H{"error:": "Creating todo"})
	}
}

func (th *TodoHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo

	if (c.BindJSON(&todo) == nil) {
		//check bool
		th.Db.Update(id, &todo)
		c.JSON(200,gin.H{"status": "Todo updated"})
	} else {
		c.JSON(500, gin.H{"error:": "Updating todo"})
	}
}

func (th *TodoHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if (th.Db.Delete(id)) {
		c.JSON(200,gin.H{"status": "Todo deleted"})
	} else {
		c.JSON(404, gin.H{"error:": "Todo not found"})
	}
}
