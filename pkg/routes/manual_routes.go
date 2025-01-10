package routes

import (
	"log"
	"net/http"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/handlers"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/mymiddleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Create_Manual_Routes(app *config.AppConfig) (_ http.Handler) {
	defer func() { //this function recovers from panic -> TODO add this function to almost all functions
		r := recover()
		if r != nil {
			//re_err = fmt.Errorf("%v", r)
		}
	}()
	mux := chi.NewRouter()
	//middleware

	mux.Use(mymiddleware.NoSurve)
	mux.Use(middleware.Recoverer)
	mux.Use(mymiddleware.Create_SessionLoad(app))
	//Get_requests

	//add generic pages (Get Requests):
	err := handlers.Add_Generic_Handlers(mux, "templates/", "./templates/layouts", "/", app)
	if err != nil {
		log.Panic(err)
	}

	//fileserver := http.FileServer(http.Dir("./static/"))
	//mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
