package render

import (
	"html/template"
	"path/filepath"
)

// Load_Template_From_Path creates a template from a path and layout_path
func Load_Template_From_Path(path string, layout_path string) (*template.Template, error) {
	ts, err := template.New(filepath.Base(path)).ParseFiles(path) //The name needs to be the same as the name of the file
	if err != nil {
		return nil, err
	}
	layouts, err := filepath.Glob(layout_path)
	if err != nil {
		return nil, err
	}
	if len(layouts) > 0 {
		ts, err = ts.ParseGlob(layout_path)
		if err != nil {
			return nil, err
		}
	}
	return ts, nil
}
