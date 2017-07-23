//Package glma lets you manage database with intuitive web-based GUI

package dbm

import(
	"fmt"
	"net/http"
	"database/sql"
	"encoding/json"
	"strings"
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
	http.HandleFunc("/graph/", RenderJSON)
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.ListenAndServe(port, nil)
}

func RenderJSON(w http.ResponseWriter, r *http.Request){
	columns, data := ShowData("SELECT file, title FROM resource")
	result := &ResultSet{Column: columns, Result: data}
	
	table_list := ShowTables()

	context := &Content{Tables: table_list, Data: result}
	jsoned, _ := json.Marshal(context)
	fmt.Fprintf(w, "%s", jsoned)
}

//GetColumn() returns one column of the resulting query
func GetColumn(row string, table string) ([]string){
	res, err := db.Query("SELECT " + row + " FROM " + table)
	if err != nil{
		panic(err)
	}

	var tmp string
	var result []string
	for res.Next(){
		res.Scan(&tmp)
		result = append(result, tmp)
	}
	return result
}

//PreprocessQuery() returns the a slice containing the columns in the query and a string containing the table
func PreprocessQuery(q string) ([]string, string){
	q = strings.TrimPrefix(q, "SELECT ")
	tmp := strings.Split(q, " FROM ")
	cols := strings.Split(tmp[0], ",")
	return cols, tmp[1]
}

//ShowData() fetches the data of SELECT query
func ShowData(q string) ([]string, [][]string){
	cols, table := PreprocessQuery(q)
	var data [][]string
	for _, val := range cols{
		tmp := GetColumn(val, table)
		data = append(data, tmp)
	}
	data = Transpose(data)
	return cols, data
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

//Transpose() rearranges the result set into much more simple arrangement
func Transpose(set [][]string) [][]string{
	var formatted [][]string
	var row []string
	var empty []string
	for i := 0; i < len(set[0]); i++{
		for ix := 0; ix < len(set); ix++{
			row = append(row, set[ix][i])
		}
		formatted = append(formatted, row)
		row = empty
	}
	return formatted
}