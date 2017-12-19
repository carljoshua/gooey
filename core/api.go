// This package starts the api, handles the requests and return json responses for managing the database

package core

import(
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

var c *Connection

// Starts the api using the args sent
func Start(db *sql.DB, driverName string) {
	c = &Connection{ database: db, driver: driverName }
	http.HandleFunc("/tables/", showTables)
	http.HandleFunc("/browse", browseTable)
	http.HandleFunc("/describe", describeTable)

	path := os.Getenv("GOPATH")
	path = path + "/src/github.com/carljoshua/gooey/client"
	http.Handle("/gooey/", http.StripPrefix("/gooey/", http.FileServer(http.Dir(path))))

	fmt.Println("Listening to port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil{
		fmt.Printf("%s\n", err)
	}
}

// showTables() returns a slice of the table names in the database
func showTables(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Content-type", "application/json")

	_, tables, err := c.GetColumn("SHOW TABLES;")
	json.NewEncoder(w).Encode(&TableList{ Tables: tables, Message: err })
}

// browseTable() returns a slice of which contains the column name and
// a two-dimensional slic that contains the data in the given table
func browseTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Content-type", "application/json")

	fields, rows, err := c.GetColumns("SELECT * FROM " + r.URL.Query().Get("table"))
	json.NewEncoder(w).Encode(&Dataset{ Columns: fields, Rows: rows, Message: err })
}


// describeTable() returns the information and structure of the given table
func describeTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Content-type", "application/json")

	fields, rows, err := c.GetColumns("DESCRIBE " + r.URL.Query().Get("table"))
	json.NewEncoder(w).Encode(&Dataset{ Columns: fields, Rows: rows, Message: err })
}
