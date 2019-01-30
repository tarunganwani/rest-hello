package model

import (
	"fmt"
	"github.com/tarunganwani/rest-hello/tdd-model-handler/errorutil"
)

type Todo struct {
	ID      int    `json:"ID"`
	Message string `json:"message"`
}

type Todolist []Todo

var todolist Todolist

type TodoModel struct {
}

func (t Todo) String() string {
	return fmt.Sprintf("{TODO id:%d, message:%s}", t.ID, t.Message)
}

func (m *TodoModel) Initialize() error {
	//todolist = []Todolist{}
	return nil
}

//CRUDs

func (m *TodoModel) CreateTodo(todoitem Todo) (Todo, error) {

	todoArrLen := len(todolist)
	newID := 0
	if todoArrLen == 0 {
		newID = 1
	} else {
		newID = todolist[todoArrLen-1].ID + 1
	}
	todoitem.ID = newID
	todolist = append(todolist, todoitem)
	return todoitem, nil
}

func (m *TodoModel) GetAllTodos() (Todolist, error) {
	if todolist == nil {
		return Todolist{}, nil
	}
	return todolist, nil
}

func (m *TodoModel) GetTodo(ID int) (todo Todo, err error) {
	err = nil
	found := false
	for _, todoItem := range todolist {
		if todoItem.ID == ID {
			found = true
			todo = todoItem
			break
		}
	}
	if found == false {
		err = errorutil.NewNotFoundError("item not found")
	}
	return
}

func (m *TodoModel) UpdateTodo(ID int, todoitemArg Todo) (Todo, error) {
	for i, todoItem := range todolist {
		if todoItem.ID == ID {
			todolist[i].Message = todoitemArg.Message
			return todolist[i], nil
		}
	}
	return Todo{}, errorutil.NewNotFoundError("item not found")
}

func (m *TodoModel) DeleteTodo(ID int) error {
	for i, todoItem := range todolist {
		if todoItem.ID == ID {
			todolist = append(todolist[:i], todolist[i+1:]...)
			return nil
		}
	}
	return errorutil.NewNotFoundError("item not found")
}

func (m *TodoModel) ClearTodolist() {
	todolist = nil
}
