package golangmyadmin

import (
	"html/template"
	"net/http"
)

// Load parse and returns the templates
func LoadUI() *template.Template{
	files, err := template.ParseGlob("UI/*.html")
	if err != nil {
		panic(err)
	}

	return files
}

//Render renders the template that have the name in the input
func Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl := LoadUI()
	tmpl.ExecuteTemplate(w, name, data)
}