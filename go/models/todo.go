package models

type Todo struct {
	Id 		int		`json:"id"`
	Name 	string	`json:"name"`
	Done	bool	`json:"done"`
}

type Todos []Todo
