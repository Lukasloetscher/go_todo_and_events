package url_variables

import (
	"fmt"
	"net/http"
)

func Get_vars(r *http.Request) (_ map[string]string, re_err error) {
	defer func() { //thsi function recovers from panic -> TODO add this function to almost all functions
		r := recover()
		if r != nil {
			re_err = fmt.Errorf("%v", r)
		}
	}()

	stringMap := make(map[string]string)
	keys := r.URL.Query()
	for key, content := range keys {
		stringMap[key] = fmt.Sprint(content)
	}

	return stringMap, nil
}
