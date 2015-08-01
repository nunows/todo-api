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
		c.JSON(404, models.Status{Success: false, Msg: "Error: Todo not found."})
	}
}

func (th *TodoHandler) Create(c *gin.Context) {
	var todo models.Todo

	if (c.BindJSON(&todo) == nil) {
		th.Db.Insert(&todo)
		c.JSON(200, models.Status{Success: true, Msg: "Todo created."})
	} else {
		c.JSON(500, models.Status{Success: false, Msg: "Error: Missing required fields."})
	}
}

func (th *TodoHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo

	if (c.BindJSON(&todo) == nil) {
		if (th.Db.Update(id, &todo)) {
			c.JSON(200, models.Status{Success: true, Msg: "Todo updated."})
		} else {
			c.JSON(500, models.Status{Success: false, Msg: "Error: Updating Todo."})
		}
	} else {
		c.JSON(500, models.Status{Success: false, Msg: "Error: Missing required fields."})
	}
}

func (th *TodoHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if (th.Db.Delete(id)) {
		c.JSON(200, models.Status{Success: true, Msg: "Todo updated."})
	} else {
		c.JSON(404, models.Status{Success: false, Msg: "Error: Todo not found."})
	}
}
