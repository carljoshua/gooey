package core

import "database/sql"

type Connection struct {
	database	*sql.DB
	driver		string
}

type Dataset struct {
	Columns		[]string	`json:"col_names"`
	Rows		[][]string	`json:"rows"`
	Message 	error		`json:"error"`
}

type TableList struct {
	Tables 		[]string	`json:"tables"`
	Message 	error		`json:"error"`
}
