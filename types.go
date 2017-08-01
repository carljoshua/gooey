package golangmyadmin

type ResultSet struct{
	Column []string	
	Result [][]string
}

type Content struct{
	Tables []string
	Data *ResultSet
	Errors map[string]string
}

