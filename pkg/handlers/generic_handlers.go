package handlers

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/config"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/models"
	"github.com/Lukasloetscher/go_todo_and_events/pkg/render"
	"github.com/go-chi/chi"
)

// Add_Generic_Handlers creates a handler for each page with the name ".generic_page.tmpl" and then adds thius to the mux.
// TODO handle layout path
func Add_Generic_Handlers(mux *chi.Mux, tmpl_path string, layout_path string, root_site string, app_ptr *config.AppConfig) error {

	err := filepath.WalkDir(tmpl_path, func(path string, d fs.DirEntry, err error) error {
		//first we check if this is indead a generic filename
		if strings.HasSuffix(path, ".generic_page.tmpl") {
			//we need to rename the file into the specified structure.
			//while doing this, we also need to save the path of the file, so we can load the document
			name, _ := strings.CutSuffix(path, ".generic_page.tmpl")
			name = strings.ReplaceAll(name, "\\", "/")
			name, _ = strings.CutPrefix(name, tmpl_path) //probably sghoudl check hgere and do a general clean up TODO
			name = root_site + name

			//We now define the handler for this and add it to the handlers_get map, the key will be the path, this makes sure, that we do not have multiple sites, that go to the same site
			mux.Get(name, func(w http.ResponseWriter, r *http.Request) {
				//Here we need to render the page...
				render.RenderTemplate(w, path, &models.TemplateData{}, app_ptr, r)
			})

			//fmt.Println("added get handler for ", name)

		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
