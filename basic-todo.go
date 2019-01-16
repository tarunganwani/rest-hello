package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Todo - Basic task structure
type Todo struct {
	ID       int       `json:"id"`
	Message  string    `json:"message"`
	Schedule time.Time `json:"time"`
}

func (t Todo) String() string {
	return fmt.Sprintf("ID : %d Message : %s \nTime %s \n", t.ID, t.Message, t.Schedule.String())
}

//Todos - todo list structure
type Todos []Todo

//Global todo list
var todolist Todos

func createToDo(msg string, sch time.Time) Todo {
	todoArrLen := len(todolist)
	newID := 0
	if todoArrLen == 0 {
		newID = 1
	} else {
		newID = todolist[todoArrLen-1].ID
	}
	return Todo{ID: newID, Message: msg, Schedule: sch}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// ------------------------------------------------------------------------
// ----------------------- handlers ---------------------------------------
// ------------------------------------------------------------------------

func baseHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "list [GET]\n")
	fmt.Fprintf(w, "create [POST]\n")
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {

	var todoItem Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todoItem); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	todoArrLen := len(todolist)
	newID := 0
	if todoArrLen == 0 {
		newID = 1
	} else {
		newID = todolist[todoArrLen-1].ID + 1
	}
	todoItem.ID = newID

	todolist = append(todolist, todoItem)
	msgStr := "Added Todo\n"
	msgStr += todoItem.String()
	msgStr += "\n"
	fmt.Fprintf(w, msgStr)
}

func listTodoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Bad request id")
		return
	}
	found := false
	for _, todoItem := range todolist {
		if todoItem.ID == todoid {
			respondWithJSON(w, http.StatusOK, todoItem)
			found = true
			break
		}
	}
	if found == false {
		respondWithError(w, http.StatusNotFound, "Item not found!")
	}
}

func listAllTodosHandler(w http.ResponseWriter, req *http.Request) {

	if todolist == nil {
		fmt.Fprintf(w, "Empty TODO list")
		return
	}

	for index, todoItem := range todolist {
		msgStr := "["
		msgStr += strconv.Itoa(index + 1)
		msgStr += "]\n"
		msgStr += todoItem.String()
		msgStr += "\n"
		fmt.Fprintf(w, msgStr)
	}
}

func updateTodoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil {
		fmt.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var decodedTodoItem Todo
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&decodedTodoItem); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request payload")
		return
	}

	for i, todoItem := range todolist {
		if todoItem.ID == todoid {
			todolist[i].Message = decodedTodoItem.Message
			todolist[i].Schedule = decodedTodoItem.Schedule
			respondWithJSON(w, http.StatusOK, todolist[i])
			return
		}
	}
	respondWithError(w, http.StatusNotFound, "Item not found")

}

func deleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Bad request id")
		return
	}
	found := false
	delElementIndex := 0
	for todoItemIndex, todoItem := range todolist {
		if todoItem.ID == todoid {
			found = true
			respondWithJSON(w, http.StatusOK, todoItem)
			delElementIndex = todoItemIndex
			break
		}
	}

	if found == false {
		respondWithError(w, http.StatusNotFound, "Item not found!")
	} else {
		todolist = append(todolist[:delElementIndex], todolist[delElementIndex+1:]...)
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", baseHandler)
	router.HandleFunc("/create", createTodoHandler).Methods("POST")
	router.HandleFunc("/todos", listAllTodosHandler).Methods("GET")
	router.HandleFunc("/todo/{todoid:[0-9]+}", listTodoHandler).Methods("GET")
	router.HandleFunc("/todo/{todoid:[0-9]+}", updateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{todoid:[0-9]+}", deleteTodoHandler).Methods("DELETE")

	fmt.Println("Listening on 8080..")
	http.ListenAndServe(":8080", router)
}
