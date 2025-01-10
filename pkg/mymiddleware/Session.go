package mymiddleware

import (
	"net/http"

	mysessions "github.com/Lukasloetscher/go_todo_and_events/pkg/MySessions"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
)

// Create_SessionLoad creates the middleware SessionLoad with a pointer to the correct app.
// SessionLoad loads and saves the session on every request
func Create_SessionLoad(app_ptr *config.AppConfig) func(next http.Handler) http.Handler {

	if (*app_ptr).Data.Session == nil {
		mysessions.Initialise_Session(app_ptr)
	}

	return func(next http.Handler) http.Handler {
		return (*app_ptr).Data.Session.LoadAndSave(next)

	}
}
