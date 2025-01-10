package config

import (
	"html/template"
	"net/smtp"

	"github.com/alexedwards/scs/v2"
)

//This should not be dependent on any other packages

// Holds the application config
type AppConfig struct {
	UseCache      bool
	ForcePreCache bool
	Portnumber    int64
	InProduction  bool
	Data          Backend_data
	Email_auth    map[string]Email_sending
}

type Backend_data struct {
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}

type Email_sending struct {
	Addr   string //addr of the server
	Sender string //for which account does this server send emails
	// maybe we should change the datastructure for this later, to allow sending from different emails
	Auth smtp.Auth //the authentication used in order to send the email
}
