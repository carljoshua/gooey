package golangmyadmin

import(
	"net/http"
	"database/sql"
	"encoding/json"
)

var (
	db *Connection
)

// Returns a mux which the api routes are added
func NewMux(database *sql.DB, db_name string) *http.ServeMux{
	db = &Connection{
		name :	db_name,
		driver : database}

	r := http.NewServeMux()
	r.HandleFunc("/db", API)

	return r
}

func API(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-type", "application/json")
	resp := &Response{}

	//get the tables in the database
	_, tables, err := db.GetColumn("SHOW TABLES;")
	if err != nil{
		resp.Errors = append(resp.Errors, err)
	}else{
		resp.Tables = tables
	}

	var q string

	if r.URL.Query().Get("query") != ""{
		q = r.URL.Query().Get("query")
	}else if r.URL.Query().Get("table") != ""{
		q = "SELECT * FROM " + r.URL.Query().Get("table")
	}

	//if the table has a value, get the structure
	if r.URL.Query().Get("table") != ""{
		//get the structure of the table
		col, res, err := db.GetColumns("DESCRIBE " + r.URL.Query().Get("table"))
		if err != nil{
			resp.Errors = append(resp.Errors, err)
		}else{
			resp.Structure = &Data{Column: col, Rows: res}
		}
	}

	if q != ""{
		//get the data from tables
		col, res, err := db.GetColumns(q)
		if err != nil{
			resp.Errors = append(resp.Errors, err)
		}else{
			resp.Dataset = &Data{Column: col, Rows: res}
		}
	}

	//convert the response into json format
	json.NewEncoder(w).Encode(resp)
}
