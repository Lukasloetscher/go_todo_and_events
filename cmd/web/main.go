package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/routes"
)

func main() {

	test()
	return
	var app_ptr *config.AppConfig
	app_ptr, err := config.Initialise_Settings()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt((*app_ptr).Portnumber, 10),
		Handler: routes.Create_Manual_Routes(app_ptr),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
