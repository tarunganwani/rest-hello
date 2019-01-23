package app

import (
	"log"

	"github.com/tarunganwani/rest/model-handler-app/model"
	"github.com/tarunganwani/rest/model-handler-app/router"
)

//App ... application structure
type App struct {
	restHandler *router.RestRouter
	Address     string
}

//InitApplication ... application init
func (a *App) InitApplication() error {

	log.Println("Initizalizing router..")
	a.restHandler = new(router.RestRouter)
	a.restHandler.Address = a.Address

	err := a.restHandler.InitRoutes()
	if err != nil {
		return err
	}

	log.Println("Initizalizing database..")
	if err = model.InitDB(); err != nil {
		return err
	}
	return nil
}

//Run ... run application
func (a *App) Run() error {
	log.Println("Listening on " + a.Address + "..")
	return a.restHandler.ListenAndServe()
}
