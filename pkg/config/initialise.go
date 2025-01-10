package config

import "html/template"

// Initialise_Settings() sets up the configuration, for the initial version.
//This returns a pointer to the newly defined Appconfig struct.
//i.e. this plays the role of a constructor
//this is mostly for testing and should be replaces by reading from a file later.
func Initialise_Settings() (*AppConfig, error) {
	var m AppConfig
	m.InProduction = false
	m.UseCache = false //only currently so i do not need to start the program new each time
	m.ForcePreCache = false
	m.Portnumber = 8080
	m.Data.TemplateCache = make(map[string]*template.Template)
	m.Email_auth = make(map[string]Email_sending)

	return &m, nil
}
