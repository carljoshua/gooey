// This package starts the api, handles the requests and return json responses for managing the database

package core

import(
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"
)

var c *Connection

// Starts the api using the args sent
func Start(db *sql.DB, driverName string, port string) {
	c = &Connection{database: db, driver: driverName}
	http.HandleFunc("/db", respond)

	fmt.Println("Listening to port " + port + "...")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil{
		fmt.Println("Invalid Port Number" + port)
	}
}

// Handles the API requests and returns json responses
func respond(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-type", "application/json")
	resp := &Response{}

	//get the tables in the database
	_, tables, err := c.GetColumn("SHOW TABLES;")
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
		col, res, err := c.GetColumns("DESCRIBE " + r.URL.Query().Get("table"))
		if err != nil{
			resp.Errors = append(resp.Errors, err)
		}else{
			resp.Structure = &Data{Column: col, Rows: res}
		}
	}

	if q != ""{
		//get the data from tables
		col, res, err := c.GetColumns(q)
		if err != nil{
			resp.Errors = append(resp.Errors, err)
		}else{
			resp.Dataset = &Data{Column: col, Rows: res}
		}
	}

	//convert the response into json format
	json.NewEncoder(w).Encode(resp)
}
