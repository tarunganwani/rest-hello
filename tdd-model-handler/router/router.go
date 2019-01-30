package router

import (
	"errors"
	"log"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tarunganwani/rest-hello/tdd-model-handler/model"
	"github.com/tarunganwani/rest-hello/tdd-model-handler/errorutil"
)

type Router struct {
	Router   *mux.Router
	model    *model.TodoModel
	initflag bool
}

func (r *Router) Initialize(modelArg *model.TodoModel) error {

	r.model = modelArg
	if r.model == nil {
		return errors.New("Router initialize error: model can not be nil")
	}

	r.Router = mux.NewRouter()
	//CRUD handlers
	r.Router.HandleFunc("/todo", r.CreateTodoHandler).Methods("POST")
	r.Router.HandleFunc("/todos", r.ListAllTodoHandler).Methods("GET")
	r.Router.HandleFunc("/todos", r.DeleteAllTodos).Methods("DELETE") // delete all todos
	r.Router.HandleFunc("/todo/{todoid:[0-9]+}", r.GetTodoHandler).Methods("GET")
	r.Router.HandleFunc("/todo/{todoid:[0-9]+}", r.UpdateTodoHandler).Methods("PUT")
	r.Router.HandleFunc("/todo/{todoid:[0-9]+}", r.DeleteTodoHandler).Methods("DELETE")

	r.initflag = true
	return nil
}

func (r *Router) Run(address string) error {

	if r.initflag == false {
		return errors.New("Must initialize before listening to requests")
	}

	log.Println("Listening on " + address)
	return http.ListenAndServe(address, r.Router)
}

func (r *Router) CreateTodoHandler(w http.ResponseWriter, req *http.Request) {


	var todoitem  model.Todo
	
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	
	err := decoder.Decode(&todoitem)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	todoitem, err = r.model.CreateTodo(todoitem)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendJsonResponse(w, http.StatusCreated, todoitem)
	
	return
}

func (r *Router) ListAllTodoHandler(w http.ResponseWriter, req *http.Request) {

	todos, err := r.model.GetAllTodos()
	if err != nil{
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	SendJsonResponse(w, http.StatusOK, todos)
	return
}

func (r *Router) GetTodoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil{
		SendErrorResponse(w, http.StatusBadRequest,  "bad request")
		return
	}
	todo, err := r.model.GetTodo(todoid)
	if err != nil{
		switch err.(type){
		case *errorutil.NotFoundError:
			SendErrorResponse(w, http.StatusNotFound, err.Error())
		default:
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	SendJsonResponse(w, http.StatusOK, todo)
	return
}

func (r *Router) UpdateTodoHandler(w http.ResponseWriter, req *http.Request) {
	
	//Get todo id
	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil{
		SendErrorResponse(w, http.StatusBadRequest,  "bad request")
		return
	}

	//Get todoitem from payload 
	var todoitem  model.Todo
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	err = decoder.Decode(&todoitem)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}


	todoitem, err = r.model.UpdateTodo(todoid, todoitem)
	if err != nil{
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	SendJsonResponse(w, http.StatusOK, todoitem)
	return
}

func (r *Router) DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	
	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil{
		SendErrorResponse(w, http.StatusBadRequest,  "bad request")
		return
	}
	err = r.model.DeleteTodo(todoid)
	if err != nil{
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	SendJsonResponse(w, http.StatusOK, map[string]string{"result" : "success"})
	return
}

func (r *Router) DeleteAllTodos(w http.ResponseWriter, req *http.Request) {
	r.model.ClearTodolist()
	SendJsonResponse(w, http.StatusOK, map[string]string{"result":"success"})
	return
}
