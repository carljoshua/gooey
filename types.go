package golangmyadmin

type ResultSet struct{
	Column []string	
	Result [][]string
}

type Content struct{
	Tables map[string]bool
	Data *ResultSet
	Errors map[string]string
}