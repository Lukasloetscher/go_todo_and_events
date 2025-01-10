package routes

import (
	"log"
	"net/http"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/handlers"
	"github.com/go-chi/chi"
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
	/*
		mux.Use(mymiddleware.NoSurve)
		mux.Use(middleware.Recoverer)
		mux.Use(mymiddleware.Create_SessionLoad(app))
		//Get_requests
		mux.Get("/", handlers.Create_Home_Handler(app))
		mux.Get("/about", handlers.Create_About_Handler(app))
		mux.Get("/test", handlers.Create_Test_Handler(app))
		mux.Get("/forms/with_get", handlers.Create_form_with_func_get_Handler(app))
		mux.Get("/forms/with_post", handlers.Create_form_with_func_post_Handler(app))
		mux.Post("/forms/with_post", handlers.Create_Post_form_with_func_post_Handler(app))
	*/
	//add generic pages (Get Requests):
	err := handlers.Add_Generic_Handlers(mux, "templates/auto/", "./templates/layouts", "/generic/", app)
	if err != nil {
		log.Panic(err)
	}

	//fileserver := http.FileServer(http.Dir("./static/"))
	//mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
