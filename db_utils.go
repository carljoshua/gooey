package golangmyadmin

import (
	"database/sql"
)

type Connection struct{
	name string	//Name of the driver
	driver *sql.DB 		//Driver of the Database
}


//GetRow() fetches a single row
func (c *Connection) GetRow(q string) (string, []string, error) {
	var row []string
	var tmp string
	res, err := c.driver.Query(q)
	if err != nil{
		return "", nil, err
	}

	col_name, _ := res.Columns()
	for res.Next(){
		res.Scan(&tmp)
		row = append(row, tmp)
	}
	return col_name[0], row, nil
}

//GetRows() fetches multiple rows
func (c *Connection) GetRows(q string) ([]string, [][]string, error){
	var rows [][]string
	res, err := c.driver.Query(q)
	if err != nil{
		return nil, nil, err
	}
	col_names, _ := res.Columns()
	for res.Next(){
		pointers := make([]interface{}, len(col_names))
		tmp := make([]string, len(col_names))
		for i, _ := range col_names{
			pointers[i] = &tmp[i]
		}
		res.Scan(pointers...)
		rows = append(rows, tmp)
	}

	return col_names, rows, nil
}

// TableQ returns the query that shows the tables in the database
func (c *Connection) TableQ() string{
	var q string
	switch c.name {
		case "sqlite3":
			q = "SELECT name FROM sqlite_master WHERE type='table';"
		case "mysql":
		 	q = "SHOW TABLES;"
	}

	return q;
}