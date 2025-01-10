package render

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/models"
)

func RenderTemplate(w http.ResponseWriter, path string, layout_path string, td *models.TemplateData, app_ptr *config.AppConfig, r *http.Request) error {

	layout_path_ext := layout_path + "/*.tmpl"
	//First we check if the serversettings want to use cached data:
	var t *template.Template
	var found bool
	var err error
	if app_ptr.UseCache { //in this case the sites should be cached and hence we try to load them from Cache
		t, found = app_ptr.Get_template_from_cache(path)
		if !found { //if we did noGetCacheMandatoryt find this entry we need to check if runtime loading is fine and when this is the case load it.
			if app_ptr.ForcePreCache {
				return errors.New("site" + path + "was not found in cache, but ForcePreCache is" + strconv.FormatBool(app_ptr.ForcePreCache))
			} else {
				//we load the site and add it to the Cache
				t, err = Load_Template_From_Path(path, layout_path_ext)
				if err != nil {
					log.Println(err)
					return err
				}
				//Then we want to add t to the Cache, so we can use it the next time, without reading the file agains
				err = app_ptr.Add_template_to_cache(path, t)
				if err != nil {
					log.Println(err)
					return err
				}
			}

		}

	} else {
		t, err = Load_Template_From_Path(path, layout_path_ext)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//now t is the correct template, we can render this template.
	//Currently we only render it, we will add dynamic data later

	buf := new(bytes.Buffer)
	err = models.AddDefaultData(td, app_ptr, r)
	if err != nil {
		log.Println(err)
		return err
	}
	err = t.Execute(buf, td)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err

	}

	return nil

}
