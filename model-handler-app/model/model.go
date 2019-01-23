package model

import (
	"errors"
	"log"
)

//Todo ... basic todo structure
type Todo struct {
	ID      uint32 `"json:"ID"`
	Message string `json:"Message"`
}

//Todos ... list - mocks our database model
type Todos []Todo

var todolist Todos

// const maxCapacity = 1000
// const initLength = 10

//InitDB ... Initialize database
func InitDB() error {
	log.Println("Initializing Todo list..")
	// todolist = make([]Todo, initLength, maxCapacity)
	return nil
}

// ******************** CRUD operations ********************

//GetTodos ... fetch todo list
func GetTodos() (Todos, error) {

	// if todolist == nil {
	// 	return nil, errors.New("Empty TODO list")
	// }
	return todolist, nil
}

//GetTodo ... fetch todo item
func GetTodo(todoid uint32) (todo Todo, err error) {
	err = nil
	found := false
	for _, todoItem := range todolist {
		if todoItem.ID == todoid {
			found = true
			todo = todoItem
			break
		}
	}
	if found == false {
		err = errors.New("item not found")
	}
	return
}

//AddTodo ... create and add a todo
func AddTodo(todoitem Todo) (Todo, error) {

	todoArrLen := len(todolist)
	newID := uint32(0)
	if todoArrLen == 0 {
		newID = 1
	} else {
		newID = todolist[todoArrLen-1].ID + 1
	}
	todoitem.ID = newID
	todolist = append(todolist, todoitem)
	return todoitem, nil
}

//UpdateTodo ... update todo item
func UpdateTodo(todoidArg uint32, todoitemArg Todo) (Todo, error) {

	for i, todoItem := range todolist {
		if todoItem.ID == todoidArg {
			todolist[i].Message = todoitemArg.Message
			return todolist[i], nil
		}
	}
	return Todo{}, errors.New("Item Not Found")
}

//DeleteTodo ... delete todo item
func DeleteTodo(todoidArg uint32) error {
	for i, todoItem := range todolist {
		if todoItem.ID == todoidArg {
			todolist = append(todolist[:i], todolist[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}
