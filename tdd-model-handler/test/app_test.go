package test

import (
	"testing"

	"github.com/tarunganwani/rest-hello/tdd-model-handler/app"
)

var a app.Application

func TestApp(t *testing.T) {
	a = app.Application{}

	hostaddress := ""
	port := "8080"
	err := a.Initialize(hostaddress, port)
	if err != nil {
		t.Errorf("Error while initializing App. Error : %s\n", err.Error())
	}
	if err = a.Run(); err != nil {
		t.Errorf("Error running the application. Error : %s\n", err.Error())
	}
}
