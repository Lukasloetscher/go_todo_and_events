package mymiddleware

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurve saves ...
func NoSurve(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		//Secure:   app.InProduction,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}
