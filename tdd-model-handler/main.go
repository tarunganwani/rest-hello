package main

import (
	"log"

	"github.com/tarunganwani/rest-hello/tdd-model-handler/app"
)

func main() {

	a := app.Application{}

	hostaddress := ""
	port := "8080"
	err := a.Initialize(hostaddress, port)

	if err != nil {
		log.Fatalf("Error while initializing App. Error : %s\n", err.Error())
	}

	if err = a.Run(); err != nil {
		log.Fatalf("Error running the application. Error : %s\n", err.Error())
	}
}
