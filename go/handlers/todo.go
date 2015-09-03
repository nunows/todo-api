package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nunows/todo-api/go/models"
)

//Common web error
var (
	invalidParamMsg = models.Status{Success: false, Msg: "Error: Invalid Parameter."}
	missingParamMsg = models.Status{Success: false, Msg: "Error: Missing required Parameter."}
	notFoundMsg     = models.Status{Success: false, Msg: "Error: Not found."}
)

type TodoHandler struct {
	Db models.TodoDb
}

func (th *TodoHandler) GetAll(c *gin.Context) {
	todos, err := th.Db.GetAll()
	if err != nil {
		c.JSON(500, models.Status{Success: false, Msg: err.Error()})
		return
	}
	c.JSON(200, todos)
}

func (th *TodoHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, invalidParamMsg)
		return
	}

	todo, err := th.Db.Get(id)
	if err != nil {
		c.JSON(404, notFoundMsg)
	} else {
		c.JSON(200, todo)
	}

}

func (th *TodoHandler) Create(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(500, missingParamMsg)
		return
	}
	err = th.Db.Insert(&todo)
	if err != nil {
		c.JSON(500, models.Status{Success: false, Msg: err.Error()})
	} else {
		c.JSON(200, models.Status{Success: true, Msg: "Todo created."})
	}
}

func (th *TodoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, invalidParamMsg)
		return
	}
	var todo models.Todo

	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(500, missingParamMsg)
		return
	}

	err = th.Db.Update(id, &todo)
	if err != nil {
		c.JSON(500, models.Status{Success: false, Msg: err.Error()})
	} else {
		c.JSON(200, models.Status{Success: true, Msg: "Todo updated."})
	}
}

func (th *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, invalidParamMsg)
		return
	}

	err = th.Db.Delete(id)
	if err != nil {
		c.JSON(404, models.Status{Success: false, Msg: err.Error()})
	} else {
		c.JSON(200, models.Status{Success: true, Msg: "Todo updated."})
	}

}
