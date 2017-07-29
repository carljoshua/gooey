//Package glma lets you manage database with intuitive web-based GUI

package dbm

import(
	//"fmt"
	"net/http"
	"database/sql"
	"strings"
	"html/template"
)

var (
	db *sql.DB
	port string
	driver string
)

type ResultSet struct{
	Column []string	
	Result [][]string
}

type Content struct{
	Tables []string
	Data *ResultSet
}

func Sync(connection *sql.DB, port_num string, driver_name string){
	db = connection
	port = port_num
	driver = driver_name
	Serve()
}

func Serve(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/golangmyadmin/", Render)
	http.ListenAndServe(port, nil)
}

func Render(w http.ResponseWriter, r *http.Request){
	table_list := ShowTables()
	context := &Content{Tables: table_list, Data: nil}
	active := r.URL.Path[len("/golangmyadmin/"):]

	result := &ResultSet{Column: nil, Result: nil}
	if r.Method == "POST" || strings.Contains(active, "browse="){
		var query string 
		if strings.Contains(active, "browse="){
			query = "SELECT COUNT(*) FROM " + active[len("browse="):]
		}else{
			query = r.FormValue("query")
		}
		columns, data := ShowData(query)
		result.Column = columns
		result.Result = data
		context.Data = result
	}
	
	var gui *template.Template
	switch active{
		case "exec":
			gui = template.Must(template.New("main").ParseFiles("base.html", "sql.html"))
		case "result":
			gui = template.Must(template.New("main").ParseFiles("base.html", "result.html"))
		default:
			if strings.Contains(active, "browse="){
				gui = template.Must(template.New("main").ParseFiles("base.html", "result.html"))
			}else{
				gui = template.Must(template.New("main").ParseFiles("base.html"))
			}
	}
	
	gui.ExecuteTemplate(w, "base", context)
}

//ShowData() fetches the data of SELECT query
func ShowData(q string ) ([]string, [][]string){
	res, err := db.Query(q)
	var rows [][]string
	if err != nil{
		panic(err)
	}
	col_names, _ := res.Columns()
	for res.Next(){
		pointers := make([]interface{}, len(col_names))
		tmp := make([]string, len(col_names))
		for i, _ := range col_names{
			pointers[i] = &tmp[i]
		}
		err := res.Scan(pointers...)
		if err != nil{
			panic(err)
		}
		rows = append(rows, tmp)
	}

	return col_names, rows
}

//ShowTables() fetches the tables inside the selected database
func ShowTables() []string {
	var tables []string
	var tmp string
	var q string
	switch driver {
		case "sqlite3":
			q = "SELECT name FROM sqlite_master WHERE type='table';"
		case "mysql":
		 	q = "SHOW TABLES"
	}
	res, err := db.Query(q)
	if err != nil{
		panic(err)
	}

	for res.Next(){
		res.Scan(&tmp)
		tables = append(tables, tmp)
	}
	return tables
}