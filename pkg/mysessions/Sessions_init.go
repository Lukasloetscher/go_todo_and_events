package mysessions

import (
	"net/http"
	"time"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
	"github.com/alexedwards/scs/v2"
)

func Initialise_Session(app_ptr *config.AppConfig) error {

	(*app_ptr).Data.Session = scs.New()
	(*app_ptr).Data.Session.Lifetime = 24 * time.Hour
	(*app_ptr).Data.Session.Cookie.Persist = true
	(*app_ptr).Data.Session.Cookie.SameSite = http.SameSiteLaxMode
	(*app_ptr).Data.Session.Cookie.Secure = (*app_ptr).InProduction

	return nil
}
