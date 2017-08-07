package golangmyadmin

import(
	"net/http"
	"database/sql"
	"strings"
)

var db *Connection

// ResultSet is a container for the rows return by the query
type ResultSet struct{
	Column []string
	Result [][]string
}

// UIValues is used to store the data that the UI rely on (e.g Table list, name of selected table)
type UIValues struct {
	List []string
	Active string
	Panel map[string]bool
	FailedQuery string
}

// Main interface where the data that will be passed into the template is stored
type Content struct{
	UI *UIValues
	Data *ResultSet
	Info []string
	Errors []error
}

// Run starts the system and serves the UI and necessary files
func Run(database *sql.DB, db_name string, port string) {
	db = &Connection{
		name :	db_name,
		driver : database}

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("UI/files"))))
	http.HandleFunc("/golangmyadmin/", Process)
	http.ListenAndServe(port, nil)
}

// Process is called everytime the website loads. It also manages requests and query execution
func Process(w http.ResponseWriter, r *http.Request){
	var tmpl string
	c := &Content{UI: &UIValues{}}

	// gets the tables in the database
	_, table_list, err := db.GetRow(db.GetQuery("table"))
	if err != nil{
		c.Errors = append(c.Errors, err)
	}else{
		c.UI.List = table_list
	}

	// fetches the selected table
	table := r.URL.Query().Get("table")
	c.UI.Active = table

	// map for highlighting the selected menu
	p := map[string]bool{"browse": false, "structure": false, "sql": false}

	// the meat. mmmm... meat!
	switch r.URL.Query().Get("do"){
		//browse tables
		case "browse":
			col, res, err := db.GetRows("SELECT * FROM " + table)
			if err != nil{
				c.Errors = append(c.Errors, err)
			}else{
				c.Data = &ResultSet{Column: col, Result: res}
			}

			p["browse"] = true
			tmpl = "result"

		//execute queries
		case "sql":
			p["sql"] = true
			tmpl = "query"

		//view the results of the query
		case "result":
			q := r.FormValue("query")
			// if query contains SELECT
			if strings.Contains(q, "SELECT"){
				col, res, err := db.GetRows(q)
				if err != nil{
					c.Errors = append(c.Errors, err)
					c.UI.FailedQuery = q
					tmpl = "query"
				}else{
					c.Data = &ResultSet{Column: col, Result: res}
					tmpl = "result"
				}
			}else{
				//If the query doesn't return rows
				err = db.Execute(q)
				if err != nil{
					c.Errors = append(c.Errors, err)
					c.UI.FailedQuery = q
				}
				tmpl = "query"
			}
			p["sql"] = true

		case "structure":

		default:
			tmpl = "main"
	}
	// storing the hightlighting params in the interface
	c.UI.Panel = p

	//renders the selected template
	Render(w, tmpl, c)
}
