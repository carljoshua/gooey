package core

import "database/sql"

type Connection struct {
	database	*sql.DB
	driver		string
}

type Response struct {
	Tables		[]string	`json:"table"`
	Structure	*Data		`json:"structure"`
	Dataset		*Data		`json:"dataset"`
	Errors		[]error		`json:"error"`
}

type Data struct {
	Column		[]string	`json:"column_name"`
	Rows		[][]string	`json:"data"`
}
