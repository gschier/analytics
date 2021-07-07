package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func RenderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) {
	if tpl == nil || !Config.CacheTemplates {
		var err error
		tpl, err = template.New("index").ParseGlob("./templates/**")
		if err != nil {
			RespondError(w, err)
			return
		}
	}

	err := tpl.ExecuteTemplate(w, name, data)
	if err != nil {
		RespondError(w, err)
	}
}
