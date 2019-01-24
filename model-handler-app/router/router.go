package router

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tarunganwani/rest/model-handler-app/errorutil"
	"github.com/tarunganwani/rest/model-handler-app/model"
)

func sendJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)

}

func sendErrorResponse(w http.ResponseWriter, code int, message string) {
	sendJSONResponse(w, code, map[string]string{"error": message})
}

//Router ..  router structure
type RestRouter struct {
	router  *mux.Router
	Address string
}

//InitRoutes .. initialize router module
func (r *RestRouter) InitRoutes() error {
	r.router = mux.NewRouter()
	r.router.HandleFunc("/todo", createTodoHandler).Methods("POST")
	r.router.HandleFunc("/todo/{todoid:[0-9]+}", listTodoHandler).Methods("GET")
	r.router.HandleFunc("/todo/{todoid:[0-9]+}", updateTodoHandler).Methods("PUT")
	r.router.HandleFunc("/todo/{todoid:[0-9]+}", deleteTodoHandler).Methods("DELETE")
	r.router.HandleFunc("/todos", listAllTodosHandler).Methods("GET")
	return nil
}

//ListenAndServe .. serve
func (r *RestRouter) ListenAndServe() error {
	return http.ListenAndServe(r.Address, r.router)
}

func createTodoHandler(w http.ResponseWriter, req *http.Request) {

	var todoitem model.Todo
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	err := decoder.Decode(&todoitem)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad payload request")
		return
	}
	todoitem, err = model.AddTodo(todoitem)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendJSONResponse(w, http.StatusOK, todoitem)
}

func listTodoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad todoid request!")
		return
	}
	todoitem, err := model.GetTodo(uint32(todoid))
	if err != nil {
		switch err.(type) {
		case *errorutil.NotFoundError:
			sendErrorResponse(w, http.StatusNotFound, err.Error())
		default:
			sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendJSONResponse(w, http.StatusOK, todoitem)
}

func listAllTodosHandler(w http.ResponseWriter, req *http.Request) {

	todolist, err := model.GetTodos()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendJSONResponse(w, http.StatusOK, todolist)
}

func updateTodoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Todo item not found")
		return
	}

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	var todoitem model.Todo
	err = decoder.Decode(&todoitem)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Bad request payload")
		return
	}

	todoitem, err = model.UpdateTodo(uint32(todoid), todoitem)
	if err != nil {
		switch err.(type) {
		case *errorutil.NotFoundError:
			sendErrorResponse(w, http.StatusNotFound, err.Error())
		default:
			sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendJSONResponse(w, http.StatusOK, todoitem)
}

func deleteTodoHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	todoid, err := strconv.Atoi(vars["todoid"])
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Todo item not found")
		return
	}

	err = model.DeleteTodo(uint32(todoid))
	if err != nil {
		switch err.(type) {
		case *errorutil.NotFoundError:
			sendErrorResponse(w, http.StatusNotFound, err.Error())
		default:
			sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendJSONResponse(w, http.StatusOK, map[string]string{"result": "success"})
}
