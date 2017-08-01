//Package golangmyadmin lets you manage database with intuitive web-based GUI

package golangmyadmin

import(
	"net/http"
	"database/sql"
	"strings"
)

var (
	db *Connection
)

// Run starts the system
func Run(database *sql.DB, db_name string, port string) {
	db = &Connection{
		name :	db_name,
		driver : database}

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("UI/files"))))
	http.HandleFunc("/golangmyadmin/", Execute)
	http.ListenAndServe(port, nil)
}


func Execute(w http.ResponseWriter, r *http.Request){
	_, table_list, _ := db.GetRow(db.TableQ())
	content := &Content{Tables: table_list}
	
	active := r.URL.Path[len("/golangmyadmin/"):]

	if r.Method == "POST" || strings.Contains(active, "browse="){
		var query string 
		if strings.Contains(active, "browse="){
			query = "SELECT * FROM " + active[len("browse="):]
		}else{
			query = r.FormValue("query")
		}
		columns, data, _ := db.GetRows(query)
		result := &ResultSet{Column: columns, Result: data}
		content.Data = result
	}
	
	switch active{
		case "exec":
			Render(w, "query", content)
		case "result":
			Render(w, "result", content)
		default:
			if strings.Contains(active, "browse="){
				Render(w, "result", content)
			}else{
				Render(w, "main", content)
			}

	}
}