package core

//GetColumns fetches a single column
func (c *Connection) GetColumn(q string) (string, []string, error) {
	var row []string
	var tmp string
	res, err := c.database.Query(q)
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

//GetColumns fetches multiple columns
func (c *Connection) GetColumns(q string) ([]string, [][]string, error){
	var rows [][]string
	res, err := c.database.Query(q)
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

//Execute runs the queries that have doesn't return rows
func (c *Connection) Execute(q string) error{
	_, err := c.database.Exec(q)
	if err != nil{
		return err
	}
	return nil
}
