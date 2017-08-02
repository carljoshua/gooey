package golangmyadmin

import (
	"html/template"
	"net/http"
	"os"
	"io/ioutil"
	"errors"
)

// BuildUI creates the files required in the user interface
func BuildUI() error {
	err := os.MkdirAll("UI/files", os.ModePerm)
	if err != nil {
		return err
	}

	for key, val := range templates {
		f, err := os.Create("./UI/" + key)
		defer f.Close()
		if err != nil{
			return err
		}

		ioutil.WriteFile("./UI/" + key, val, 0600)
	}

	for key, val := range files {
		f, err := os.Create("./UI/files/" + key)
		defer f.Close()
		if err != nil{
			return err
		}

		ioutil.WriteFile("./UI/files/" + key, val, 0600)
	}

	return nil
}

// Load parse and returns the templates
func LoadUI() (*template.Template, error){
	files, err := template.ParseGlob("UI/*.html")
	if err != nil {
		if _, errs := os.Stat("UI"); os.IsNotExist(errs){
			return nil, errors.New("Build UI first")
		}else{
			return nil, err
		}
	}

	return files, nil
}

//Render renders the template that have the name in the input
func Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := LoadUI()
	if err != nil{
		if err.Error() == "Build UI first"{
			BuildUI()
			tmpl, _ = LoadUI()
		}else{
			panic(err)
		}
	}

	tmpl.ExecuteTemplate(w, name, data)
}