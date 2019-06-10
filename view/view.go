package view

import (
	"html/template"
	"os"

	. "github.com/mitchdennett/flameframework"
)

func Render(templatePath string, data interface{}) {
	d, err := os.Getwd()
	if err != nil {
		return
	}

	w := Current.GetResponse()

	t := template.Must(template.ParseGlob(d + "/templates/*"))
	t.ExecuteTemplate(w, templatePath, data)
}
