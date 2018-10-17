package view

import (
	"fmt"
	"os"
	"html/template"
	."github.com/flame"
)

func Render(templatePath string, data interface{}) {
	d, err := os.Getwd()
	fmt.Println(d)
	if err != nil {
		return
	}
	
	w := Current.GetResponse()
	
	t, _ := template.ParseFiles(d + "/templates/" + templatePath)
	t.Execute(w, data)
}