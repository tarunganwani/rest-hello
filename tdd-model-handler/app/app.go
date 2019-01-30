package app

import (
	"errors"

	"github.com/tarunganwani/rest-hello/tdd-model-handler/model"
	"github.com/tarunganwani/rest-hello/tdd-model-handler/router"
)

//Application - strcture to hold application level components
type Application struct {
	RestRouter *router.Router
	Model      *model.TodoModel
	address    string
}

func (a *Application) Initialize(ipaddress string, port string) error {

	a.address = ipaddress + ":" + port
	a.Model = new(model.TodoModel)
	a.RestRouter = new(router.Router)
	err := a.Model.Initialize()
	if err != nil {
		return errors.New("Application init error : " + err.Error())
	}
	err = a.RestRouter.Initialize(a.Model)
	if err != nil {
		return errors.New("Application init error : " + err.Error())
	}
	return nil
}

func (a *Application) Run() error {

	err := a.RestRouter.Run(a.address)
	if err != nil {
		return errors.New("Application run error : " + err.Error())
	}
	return nil
}
