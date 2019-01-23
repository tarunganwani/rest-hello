package main

import (
	"log"

	"github.com/tarunganwani/rest/model-handler-app/app"
)

func main() {
	appObj := app.App{Address: ":8080"}
	err := appObj.InitApplication()
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(appObj.Run())
}
