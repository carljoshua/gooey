package golangmyadmin

import(
	"net/http"
	"database/sql"
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
	var tmpl string
	table_list := make(map[string]bool)

	_, tables, _ := db.GetRow(db.TableQ())
	for _, name := range tables{
		table_list[name] = false
	}


	c := &Content{Tables: table_list}
	
	switch r.URL.Query().Get("do"){
		case "browse":
			table := r.URL.Query().Get("table")
			c.Tables[table] = true
			col, res, _ := db.GetRows("SELECT * FROM " + table)
			c.Data = &ResultSet{Column: col, Result: res}
			tmpl = "result"
		case "sql":
			tmpl = "query"
		case "result":
			col, res, _ := db.GetRows(r.FormValue("query"))
			c.Data = &ResultSet{Column: col, Result: res}
			tmpl = "result"
		default:
			tmpl = "main"
	}

	Render(w, tmpl, c)
}