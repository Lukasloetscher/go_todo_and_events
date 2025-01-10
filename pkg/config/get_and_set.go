package config

import (
	"html/template"
	"log"
)

func (m *AppConfig) Get_template_from_cache(key string) (*template.Template, bool) {
	temp, ok := m.Data.TemplateCache[key]
	if ok {
		return temp, true
	} else {
		return nil, false
	}
}

func (m *AppConfig) Add_template_to_cache(key string, data *template.Template) error {
	_, ok := m.Data.TemplateCache[key]
	if ok {
		log.Println("Warning: the key", key, "already exists in cached_sites and will be overwritten")
	}
	m.Data.TemplateCache[key] = data
	return nil

}
