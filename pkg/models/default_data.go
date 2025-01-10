package models

import (
	"net/http"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
	"github.com/justinas/nosurf"
)

// AddDefaultData sets default parameters, to a certain values
// Note that it might make sense, that this function gets access to the sessions.
func AddDefaultData(td *TemplateData, app_ptr *config.AppConfig, r *http.Request) error {
	if td.BoolMap == nil {
		//if the map does not yet exists, we create it
		td.BoolMap = make(map[string]bool)
	}
	_, found := td.BoolMap["navbar"]
	if !found {
		td.BoolMap["navbar"] = true //if nothing is specified, we always want to show the navbar
	}

	td.CSRFToken = nosurf.Token(r)

	return nil
}
