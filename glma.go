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
	http.HandleFunc("/golangmyadmin/", Process)
	http.ListenAndServe(port, nil)
}


func Process(w http.ResponseWriter, r *http.Request){
	var tmpl string
	table_list := make(map[string]string)

	_, tables, _ := db.GetRow(db.TableQ())
	for _, name := range tables{
		table_list[name] = ""
	}


	c := &Content{Tables: table_list}
	
	switch r.URL.Query().Get("do"){
		case "browse":
			table := r.URL.Query().Get("table")
			c.Tables[table] = "active"
			col, res, _ := db.GetRows("SELECT * FROM " + table)
			c.Data = &ResultSet{Column: col, Result: res}
			tmpl = "result"
		case "sql":
			tmpl = "query"
		case "result":
			q := r.FormValue("query")
			if strings.Contains(q, "SELECT"){
				col, res, _ := db.GetRows(q)
				c.Data = &ResultSet{Column: col, Result: res}
			}else{
				db.Execute(q)
			}
			
			tmpl = "result"
		default:
			tmpl = "main"
	}

	Render(w, tmpl, c)
}