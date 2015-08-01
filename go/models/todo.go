package models

type Todo struct {
	Id 	int	`json:"id"`
	Name 	string	`json:"name"	binding:"required"`
	Done	bool	`json:"done"	binding:"required"`
}

type Todos []Todo
