package golangmyadmin

var templates = map[string][]byte{
	"base.html"		: []byte{123,123,100,101,102,105,110,101,32,34,98,97,115,101,45,116,111,112,34,125,125,10,60,33,68,79,67,84,89,80,69,32,104,116,109,108,62,10,60,104,116,109,108,62,10,60,104,101,97,100,62,10,9,60,116,105,116,108,101,62,103,111,108,97,110,103,109,121,97,100,109,105,110,60,47,116,105,116,108,101,62,10,9,60,108,105,110,107,32,104,114,101,102,61,34,47,102,105,108,101,115,47,115,116,121,108,101,46,99,115,115,34,32,114,101,108,61,34,115,116,121,108,101,115,104,101,101,116,34,62,10,60,47,104,101,97,100,62,10,60,98,111,100,121,62,10,9,60,100,105,118,62,10,9,9,60,110,97,118,32,99,108,97,115,115,61,34,104,101,97,100,101,114,34,32,105,100,61,34,104,101,97,100,101,114,34,62,10,9,9,9,60,100,105,118,62,10,9,9,9,9,60,98,117,116,116,111,110,32,105,100,61,34,115,105,100,101,98,97,114,45,116,111,103,103,108,101,34,32,111,110,99,108,105,99,107,61,34,115,105,100,101,72,105,100,101,40,41,34,62,10,9,9,9,9,9,60,115,112,97,110,32,99,108,97,115,115,61,34,104,97,109,98,117,114,103,101,114,34,62,60,47,115,112,97,110,62,10,9,9,9,9,60,47,98,117,116,116,111,110,62,10,9,9,9,60,47,100,105,118,62,10,9,9,60,47,110,97,118,62,10,9,9,60,33,45,45,32,83,105,100,101,66,97,114,32,45,45,62,10,9,9,60,100,105,118,32,99,108,97,115,115,61,34,115,105,100,101,98,97,114,34,32,105,100,61,34,115,105,100,101,98,97,114,45,119,114,97,112,34,62,10,9,9,9,60,100,105,118,32,99,108,97,115,115,61,34,112,114,111,102,105,108,101,34,62,60,104,51,62,65,68,77,73,78,60,47,104,51,62,60,47,100,105,118,62,10,9,9,9,60,100,105,118,32,99,108,97,115,115,61,34,115,105,100,101,98,97,114,45,109,101,110,117,34,62,60,104,114,32,47,62,10,9,9,9,9,60,100,105,118,62,10,9,9,9,9,9,60,104,52,62,84,97,98,108,101,115,60,47,104,52,62,10,9,9,9,9,9,60,117,108,62,10,9,9,9,9,9,9,123,123,114,97,110,103,101,32,46,76,105,115,116,125,125,10,9,9,9,9,9,9,9,60,108,105,62,60,97,32,104,114,101,102,61,123,123,112,114,105,110,116,102,32,34,116,97,115,107,63,100,111,61,98,114,111,119,115,101,38,116,97,98,108,101,61,37,115,34,32,46,125,125,62,123,123,46,125,125,60,47,97,62,60,47,108,105,62,10,9,9,9,9,9,9,123,123,101,110,100,125,125,10,9,9,9,9,9,60,47,117,108,62,10,9,9,9,9,60,47,100,105,118,62,60,104,114,32,47,62,10,9,9,9,60,47,100,105,118,62,10,9,9,60,47,100,105,118,62,10,9,9,60,33,45,45,32,77,97,105,110,32,67,111,110,116,101,110,116,45,45,62,10,9,9,60,100,105,118,32,99,108,97,115,115,61,34,109,97,105,110,45,99,111,110,116,101,110,116,34,32,105,100,61,34,99,111,110,116,101,110,116,45,119,114,97,112,34,62,10,9,9,9,60,100,105,118,32,99,108,97,115,115,61,34,112,97,110,101,108,34,62,10,9,9,9,9,60,100,105,118,32,99,108,97,115,115,61,34,98,114,101,97,100,99,114,117,109,98,115,34,62,60,104,50,62,123,123,46,65,99,116,105,118,101,125,125,60,47,104,50,62,60,47,100,105,118,62,10,9,9,9,9,60,100,105,118,32,99,108,97,115,115,61,34,110,97,118,105,103,97,116,105,111,110,34,62,10,9,9,9,9,9,123,123,36,97,99,116,105,118,101,32,58,61,32,46,65,99,116,105,118,101,125,125,10,9,9,9,9,9,123,123,105,102,32,36,97,99,116,105,118,101,125,125,10,9,9,9,9,9,9,60,117,108,62,10,9,9,9,9,9,9,9,123,123,114,97,110,103,101,32,36,107,101,121,44,32,36,118,97,108,32,58,61,32,46,80,97,110,101,108,125,125,10,9,9,9,9,9,9,9,9,123,123,105,102,32,36,118,97,108,125,125,10,9,9,9,9,9,9,9,9,9,60,108,105,32,99,108,97,115,115,61,34,97,99,116,105,118,101,34,62,10,9,9,9,9,9,9,9,9,123,123,101,108,115,101,125,125,10,9,9,9,9,9,9,9,9,9,60,108,105,62,10,9,9,9,9,9,9,9,9,123,123,101,110,100,125,125,10,9,9,9,9,9,9,9,9,60,97,32,104,114,101,102,61,123,123,112,114,105,110,116,102,32,34,116,97,115,107,63,100,111,61,37,115,38,116,97,98,108,101,61,37,115,34,32,36,107,101,121,32,36,97,99,116,105,118,101,125,125,62,123,123,36,107,101,121,125,125,60,47,97,62,60,47,108,105,62,10,9,9,9,9,9,9,9,123,123,101,110,100,125,125,10,9,9,9,9,9,9,60,47,117,108,62,10,9,9,9,9,9,123,123,101,110,100,125,125,10,9,9,9,9,60,47,100,105,118,62,10,9,9,9,60,47,100,105,118,62,10,9,9,9,60,100,105,118,32,99,108,97,115,115,61,34,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,34,62,10,9,9,9,9,60,100,105,118,32,99,108,97,115,115,61,34,105,110,116,101,114,97,99,116,105,118,101,34,62,10,123,123,101,110,100,125,125,10,10,123,123,100,101,102,105,110,101,32,34,98,97,115,101,45,98,111,116,116,111,109,34,125,125,10,9,9,9,9,60,47,100,105,118,62,10,9,9,9,60,47,100,105,118,62,10,9,9,60,47,100,105,118,62,10,9,60,47,100,105,118,62,10,60,47,98,111,100,121,62,10,60,115,99,114,105,112,116,32,115,114,99,61,34,47,102,105,108,101,115,47,115,116,121,108,101,46,106,115,34,62,60,47,115,99,114,105,112,116,62,10,60,47,104,116,109,108,62,10,123,123,101,110,100,125,125,10},
	"query.html"	:	[]byte{123,123,100,101,102,105,110,101,32,34,113,117,101,114,121,34,125,125,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,116,111,112,34,32,46,85,73,125,125,10,123,123,119,105,116,104,32,46,85,73,46,69,114,114,111,114,115,125,125,10,32,32,32,32,60,100,105,118,32,99,108,97,115,115,61,34,101,114,114,111,114,95,112,111,112,117,112,34,62,10,32,32,32,32,32,32,32,32,123,123,114,97,110,103,101,32,46,125,125,10,32,32,32,32,32,32,32,32,32,32,32,32,123,123,112,114,105,110,116,102,32,34,37,115,34,32,46,125,125,60,98,114,32,47,62,10,32,32,32,32,32,32,32,32,123,123,101,110,100,125,125,10,32,32,32,32,60,47,100,105,118,62,10,123,123,101,110,100,125,125,10,123,123,36,97,99,116,105,118,101,32,58,61,32,46,85,73,46,65,99,116,105,118,101,125,125,10,123,123,36,102,113,32,58,61,32,46,85,73,46,70,97,105,108,101,100,81,117,101,114,121,125,125,10,60,102,111,114,109,32,97,99,116,105,111,110,61,123,123,112,114,105,110,116,102,32,34,116,97,115,107,63,100,111,61,114,101,115,117,108,116,38,116,97,98,108,101,61,37,115,34,32,36,97,99,116,105,118,101,125,125,32,109,101,116,104,111,100,61,34,80,79,83,84,34,62,10,9,60,116,101,120,116,97,114,101,97,32,110,97,109,101,61,34,113,117,101,114,121,34,32,112,108,97,99,101,104,111,108,100,101,114,61,34,81,117,101,114,105,101,115,32,72,101,114,101,34,62,123,123,105,102,32,36,102,113,125,125,123,123,36,102,113,125,125,123,123,101,110,100,125,125,60,47,116,101,120,116,97,114,101,97,62,10,9,60,105,110,112,117,116,32,116,121,112,101,61,34,115,117,98,109,105,116,34,32,118,97,108,117,101,61,34,69,120,101,99,117,116,101,34,32,99,108,97,115,115,61,34,98,116,110,45,115,117,98,109,105,116,34,62,10,60,47,100,105,118,62,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,98,111,116,116,111,109,34,32,46,125,125,10,123,123,101,110,100,125,125,10,10,123,123,100,101,102,105,110,101,32,34,114,101,115,117,108,116,34,125,125,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,116,111,112,34,32,46,85,73,125,125,10,60,116,97,98,108,101,62,10,9,60,116,114,62,10,9,123,123,114,97,110,103,101,32,46,68,97,116,97,46,67,111,108,117,109,110,125,125,10,9,9,60,116,104,62,60,104,53,62,123,123,46,125,125,60,47,104,53,62,60,47,116,104,62,10,9,123,123,101,110,100,125,125,10,9,60,47,116,114,62,10,10,9,123,123,114,97,110,103,101,32,46,68,97,116,97,46,82,101,115,117,108,116,125,125,10,9,9,60,116,114,62,10,9,9,9,123,123,114,97,110,103,101,32,46,125,125,10,9,9,9,9,60,116,100,62,123,123,46,125,125,60,47,116,100,62,10,9,9,9,123,123,101,110,100,125,125,10,9,9,60,47,116,114,62,10,9,123,123,101,110,100,125,125,10,60,47,116,97,98,108,101,62,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,98,111,116,116,111,109,34,32,46,125,125,10,123,123,101,110,100,125,125,10},
	"struct.html"	: []byte{123,123,100,101,102,105,110,101,32,34,115,116,114,117,99,116,117,114,101,34,125,125,10,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,116,111,112,34,32,46,85,73,125,125,10,10,60,116,97,98,108,101,62,10,9,60,116,114,62,10,9,123,123,114,97,110,103,101,32,46,68,97,116,97,46,67,111,108,117,109,110,125,125,10,9,9,60,116,104,62,60,104,53,62,123,123,46,125,125,60,47,104,53,62,60,47,116,104,62,10,9,123,123,101,110,100,125,125,10,9,60,47,116,114,62,10,10,9,123,123,114,97,110,103,101,32,46,68,97,116,97,46,82,101,115,117,108,116,125,125,10,9,9,60,116,114,62,10,9,9,9,123,123,114,97,110,103,101,32,46,125,125,10,9,9,9,9,60,116,100,62,123,123,46,125,125,60,47,116,100,62,10,9,9,9,123,123,101,110,100,125,125,10,9,9,60,47,116,114,62,10,9,123,123,101,110,100,125,125,10,60,47,116,97,98,108,101,62,10,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,98,111,116,116,111,109,34,125,125,10,123,123,101,110,100,125,125,10},
	"main.html"		: []byte{123,123,100,101,102,105,110,101,32,34,109,97,105,110,34,125,125,10,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,116,111,112,34,32,46,85,73,125,125,10,123,123,119,105,116,104,32,46,69,114,114,111,114,115,125,125,10,32,32,32,32,123,123,114,97,110,103,101,32,46,125,125,10,32,32,32,32,32,32,32,32,123,123,112,114,105,110,116,102,32,34,37,115,60,98,114,32,47,62,34,32,46,125,125,10,32,32,32,32,123,123,101,110,100,125,125,10,123,123,101,110,100,125,125,10,10,60,104,49,62,84,101,115,116,105,110,103,60,47,104,49,62,10,10,123,123,116,101,109,112,108,97,116,101,32,34,98,97,115,101,45,98,111,116,116,111,109,34,32,46,125,125,10,10,123,123,101,110,100,125,125,10}}

var files = map[string][]byte{
	"style.js"		:	[]byte{102,117,110,99,116,105,111,110,32,95,40,105,100,41,123,10,9,114,101,116,117,114,110,32,100,111,99,117,109,101,110,116,46,103,101,116,69,108,101,109,101,110,116,66,121,73,100,40,105,100,41,59,10,125,10,10,102,117,110,99,116,105,111,110,32,115,105,100,101,83,104,111,119,40,41,123,10,9,115,105,100,101,98,97,114,32,61,32,95,40,39,115,105,100,101,98,97,114,45,119,114,97,112,39,41,59,10,9,99,111,110,116,101,110,116,32,61,32,95,40,39,99,111,110,116,101,110,116,45,119,114,97,112,39,41,59,10,9,115,105,100,101,98,97,114,46,115,101,116,65,116,116,114,105,98,117,116,101,40,34,115,116,121,108,101,34,44,32,34,34,41,59,10,9,99,111,110,116,101,110,116,46,115,101,116,65,116,116,114,105,98,117,116,101,40,34,115,116,121,108,101,34,44,32,34,34,41,59,10,9,95,40,39,115,105,100,101,98,97,114,45,116,111,103,103,108,101,39,41,46,115,101,116,65,116,116,114,105,98,117,116,101,40,39,111,110,99,108,105,99,107,39,44,32,39,115,105,100,101,72,105,100,101,40,41,39,41,59,10,9,115,101,115,115,105,111,110,83,116,111,114,97,103,101,46,115,101,116,73,116,101,109,40,34,115,105,100,101,98,97,114,34,44,32,34,111,112,101,110,34,41,59,10,9,114,101,115,105,122,101,83,105,100,101,98,97,114,40,41,10,125,10,10,102,117,110,99,116,105,111,110,32,115,105,100,101,72,105,100,101,40,41,123,10,9,115,105,100,101,98,97,114,32,61,32,95,40,39,115,105,100,101,98,97,114,45,119,114,97,112,39,41,59,10,9,99,111,110,116,101,110,116,32,61,32,95,40,39,99,111,110,116,101,110,116,45,119,114,97,112,39,41,59,10,9,115,105,100,101,98,97,114,46,115,101,116,65,116,116,114,105,98,117,116,101,40,34,115,116,121,108,101,34,44,32,34,119,105,100,116,104,58,32,48,34,41,59,10,9,99,111,110,116,101,110,116,46,115,101,116,65,116,116,114,105,98,117,116,101,40,34,115,116,121,108,101,34,44,32,34,119,105,100,116,104,58,32,49,48,48,37,59,32,109,97,114,103,105,110,45,108,101,102,116,58,32,48,59,34,41,59,10,9,95,40,39,115,105,100,101,98,97,114,45,116,111,103,103,108,101,39,41,46,115,101,116,65,116,116,114,105,98,117,116,101,40,39,111,110,99,108,105,99,107,39,44,32,39,115,105,100,101,83,104,111,119,40,41,39,41,59,10,9,115,101,115,115,105,111,110,83,116,111,114,97,103,101,46,115,101,116,73,116,101,109,40,34,115,105,100,101,98,97,114,34,44,32,34,99,108,111,115,101,34,41,10,9,114,101,115,105,122,101,83,105,100,101,98,97,114,40,41,10,125,10,10,102,117,110,99,116,105,111,110,32,97,112,112,108,121,84,114,97,110,115,105,116,105,111,110,40,41,123,10,9,95,40,34,115,105,100,101,98,97,114,45,119,114,97,112,34,41,46,115,101,116,65,116,116,114,105,98,117,116,101,40,34,99,108,97,115,115,34,44,32,34,115,105,100,101,98,97,114,32,115,105,100,101,45,116,114,97,110,115,105,116,105,111,110,34,41,59,10,9,95,40,34,99,111,110,116,101,110,116,45,119,114,97,112,34,41,46,115,101,116,65,116,116,114,105,98,117,116,101,40,34,99,108,97,115,115,34,44,32,34,109,97,105,110,45,99,111,110,116,101,110,116,32,109,97,105,110,45,116,114,97,110,115,105,116,105,111,110,34,41,10,125,10,10,102,117,110,99,116,105,111,110,32,114,101,115,105,122,101,83,105,100,101,98,97,114,40,41,123,10,9,105,102,40,95,40,34,99,111,110,116,101,110,116,45,119,114,97,112,34,41,46,111,102,102,115,101,116,72,101,105,103,104,116,32,62,61,32,100,111,99,117,109,101,110,116,46,98,111,100,121,46,115,99,114,111,108,108,72,101,105,103,104,116,32,45,32,95,40,34,104,101,97,100,101,114,34,41,46,111,102,102,115,101,116,72,101,105,103,104,116,41,123,10,9,9,95,40,34,115,105,100,101,98,97,114,45,119,114,97,112,34,41,46,115,116,121,108,101,46,104,101,105,103,104,116,32,61,32,95,40,34,99,111,110,116,101,110,116,45,119,114,97,112,34,41,46,111,102,102,115,101,116,72,101,105,103,104,116,32,43,32,34,112,120,34,10,9,125,101,108,115,101,123,10,9,9,104,101,97,100,101,114,72,101,105,103,104,116,32,61,32,95,40,34,104,101,97,100,101,114,34,41,46,111,102,102,115,101,116,72,101,105,103,104,116,10,9,9,115,105,122,101,32,61,32,104,101,97,100,101,114,72,101,105,103,104,116,32,47,32,100,111,99,117,109,101,110,116,46,98,111,100,121,46,115,99,114,111,108,108,72,101,105,103,104,116,10,9,9,95,40,34,115,105,100,101,98,97,114,45,119,114,97,112,34,41,46,115,116,121,108,101,46,104,101,105,103,104,116,32,61,32,49,48,48,32,45,32,40,115,105,122,101,32,42,32,49,48,48,41,32,43,32,34,37,34,10,9,125,10,125,10,10,102,117,110,99,116,105,111,110,32,109,97,110,97,103,101,40,41,123,10,9,118,97,114,32,99,114,97,112,32,61,32,115,101,115,115,105,111,110,83,116,111,114,97,103,101,46,103,101,116,73,116,101,109,40,34,115,105,100,101,98,97,114,34,41,59,10,9,105,102,40,99,114,97,112,32,61,61,32,34,111,112,101,110,34,41,123,10,9,9,115,105,100,101,83,104,111,119,40,41,10,9,125,101,108,115,101,32,105,102,40,99,114,97,112,32,61,61,32,34,99,108,111,115,101,34,41,123,10,9,9,115,105,100,101,72,105,100,101,40,41,10,9,125,10,9,97,112,112,108,121,84,114,97,110,115,105,116,105,111,110,40,41,10,9,114,101,115,105,122,101,83,105,100,101,98,97,114,40,41,10,125,10,10,102,117,110,99,116,105,111,110,32,115,101,110,100,81,117,101,114,121,40,41,123,10,9,97,108,101,114,116,40,95,40,34,113,117,101,114,121,34,41,46,116,101,120,116,67,111,110,116,101,110,116,41,10,125,10,10,119,105,110,100,111,119,46,111,110,108,111,97,100,32,61,32,109,97,110,97,103,101,40,41,10},
	"style.css"		:	[]byte{42,32,123,10,32,32,102,111,110,116,45,102,97,109,105,108,121,58,32,83,97,110,115,45,115,101,114,105,102,44,32,86,101,114,100,97,110,97,44,32,84,97,104,111,109,97,59,10,32,32,112,97,100,100,105,110,103,58,32,48,112,120,59,10,32,32,109,97,114,103,105,110,58,32,48,112,120,59,10,32,32,102,111,110,116,45,119,101,105,103,104,116,58,32,50,48,48,59,32,125,10,10,46,98,116,110,44,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,102,111,114,109,32,46,98,116,110,45,115,117,98,109,105,116,32,123,10,32,32,109,97,114,103,105,110,45,116,111,112,58,32,49,51,112,120,59,10,32,32,98,111,114,100,101,114,58,32,110,111,110,101,59,10,32,32,98,97,99,107,103,114,111,117,110,100,58,32,108,105,103,104,116,99,111,114,97,108,59,10,32,32,99,111,108,111,114,58,32,119,104,105,116,101,59,10,32,32,116,101,120,116,45,116,114,97,110,115,102,111,114,109,58,32,117,112,112,101,114,99,97,115,101,59,10,32,32,99,117,114,115,111,114,58,32,112,111,105,110,116,101,114,59,10,32,32,112,97,100,100,105,110,103,58,32,49,53,112,120,59,32,125,10,10,46,104,101,97,100,101,114,32,123,10,32,32,98,97,99,107,103,114,111,117,110,100,58,32,35,51,51,51,51,51,51,59,10,32,32,112,97,100,100,105,110,103,58,32,49,53,112,120,32,49,53,112,120,32,49,56,112,120,32,50,51,112,120,59,10,32,32,99,111,108,111,114,58,32,119,104,105,116,101,59,32,125,10,32,32,46,104,101,97,100,101,114,32,100,105,118,32,123,10,32,32,32,32,100,105,115,112,108,97,121,58,32,105,110,108,105,110,101,45,98,108,111,99,107,59,32,125,10,10,35,115,105,100,101,98,97,114,45,116,111,103,103,108,101,32,123,10,32,32,112,97,100,100,105,110,103,58,32,49,50,112,120,32,49,48,112,120,32,49,48,112,120,32,49,50,112,120,59,10,32,32,98,111,114,100,101,114,58,32,48,59,10,32,32,98,97,99,107,103,114,111,117,110,100,58,32,35,55,70,56,67,56,68,59,10,32,32,99,117,114,115,111,114,58,32,112,111,105,110,116,101,114,59,10,32,32,98,111,120,45,115,104,97,100,111,119,58,32,35,54,101,54,101,54,101,32,48,112,120,32,45,51,112,120,32,50,112,120,32,105,110,115,101,116,59,32,125,10,10,46,104,97,109,98,117,114,103,101,114,32,123,10,32,32,112,111,115,105,116,105,111,110,58,32,114,101,108,97,116,105,118,101,59,10,32,32,100,105,115,112,108,97,121,58,32,105,110,108,105,110,101,45,98,108,111,99,107,59,10,32,32,119,105,100,116,104,58,32,49,46,50,53,101,109,59,10,32,32,104,101,105,103,104,116,58,32,48,46,54,51,101,109,59,10,32,32,109,97,114,103,105,110,45,114,105,103,104,116,58,32,48,46,51,101,109,59,10,32,32,98,111,114,100,101,114,45,116,111,112,58,32,48,46,50,101,109,32,115,111,108,105,100,32,35,102,102,102,59,10,32,32,98,111,114,100,101,114,45,98,111,116,116,111,109,58,32,48,46,50,101,109,32,115,111,108,105,100,32,35,102,102,102,59,32,125,10,10,46,104,97,109,98,117,114,103,101,114,58,98,101,102,111,114,101,32,123,10,32,32,99,111,110,116,101,110,116,58,32,34,34,59,10,32,32,112,111,115,105,116,105,111,110,58,32,97,98,115,111,108,117,116,101,59,10,32,32,116,111,112,58,32,48,46,50,101,109,59,10,32,32,108,101,102,116,58,32,48,112,120,59,10,32,32,119,105,100,116,104,58,32,49,48,48,37,59,10,32,32,98,111,114,100,101,114,45,116,111,112,58,32,48,46,50,101,109,32,115,111,108,105,100,32,35,102,102,102,59,32,125,10,10,47,42,32,83,105,100,101,98,97,114,32,42,47,10,46,115,105,100,101,98,97,114,32,123,10,32,32,122,45,105,110,100,101,120,58,32,49,59,10,32,32,112,111,115,105,116,105,111,110,58,32,97,98,115,111,108,117,116,101,59,10,32,32,119,105,100,116,104,58,32,50,53,37,59,10,32,32,111,118,101,114,102,108,111,119,58,32,104,105,100,100,101,110,59,10,32,32,98,97,99,107,103,114,111,117,110,100,58,32,35,49,101,51,53,52,48,59,32,125,10,32,32,46,115,105,100,101,98,97,114,32,46,112,114,111,102,105,108,101,32,123,10,32,32,32,32,98,97,99,107,103,114,111,117,110,100,58,32,35,49,57,50,100,51,55,59,10,32,32,32,32,116,101,120,116,45,97,108,105,103,110,58,32,99,101,110,116,101,114,59,10,32,32,32,32,112,97,100,100,105,110,103,58,32,50,48,112,120,32,48,59,10,32,32,32,32,100,105,115,112,108,97,121,58,32,98,108,111,99,107,59,32,125,10,32,32,32,32,46,115,105,100,101,98,97,114,32,46,112,114,111,102,105,108,101,32,104,51,32,123,10,32,32,32,32,32,32,99,111,108,111,114,58,32,119,104,105,116,101,59,32,125,10,32,32,46,115,105,100,101,98,97,114,32,46,115,105,100,101,98,97,114,45,109,101,110,117,32,100,105,118,32,104,52,32,123,10,32,32,32,32,116,101,120,116,45,116,114,97,110,115,102,111,114,109,58,32,117,112,112,101,114,99,97,115,101,59,10,32,32,32,32,99,111,108,111,114,58,32,35,55,101,56,99,57,51,59,10,32,32,32,32,112,97,100,100,105,110,103,58,32,50,48,112,120,32,48,32,49,48,112,120,32,49,48,37,59,10,32,32,32,32,102,111,110,116,45,115,105,122,101,58,32,48,46,56,53,101,109,59,32,125,10,32,32,46,115,105,100,101,98,97,114,32,46,115,105,100,101,98,97,114,45,109,101,110,117,32,100,105,118,32,117,108,32,123,10,32,32,32,32,108,105,115,116,45,115,116,121,108,101,58,32,110,111,110,101,59,10,32,32,32,32,112,97,100,100,105,110,103,45,98,111,116,116,111,109,58,32,50,48,112,120,59,32,125,10,32,32,32,32,46,115,105,100,101,98,97,114,32,46,115,105,100,101,98,97,114,45,109,101,110,117,32,100,105,118,32,117,108,32,108,105,32,123,10,32,32,32,32,32,32,100,105,115,112,108,97,121,58,32,98,108,111,99,107,59,32,125,10,32,32,32,32,32,32,46,115,105,100,101,98,97,114,32,46,115,105,100,101,98,97,114,45,109,101,110,117,32,100,105,118,32,117,108,32,108,105,32,97,32,123,10,32,32,32,32,32,32,32,32,116,101,120,116,45,100,101,99,111,114,97,116,105,111,110,58,32,110,111,110,101,59,10,32,32,32,32,32,32,32,32,99,111,108,111,114,58,32,35,57,48,57,102,97,53,59,10,32,32,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,49,48,112,120,32,48,32,49,48,112,120,32,49,53,37,59,10,32,32,32,32,32,32,32,32,100,105,115,112,108,97,121,58,32,98,108,111,99,107,59,32,125,10,32,32,46,115,105,100,101,98,97,114,32,46,115,105,100,101,98,97,114,45,109,101,110,117,32,104,114,32,123,10,32,32,32,32,98,111,114,100,101,114,58,32,48,46,53,112,120,32,115,111,108,105,100,32,35,49,97,50,100,51,55,59,32,125,10,10,47,42,32,77,97,105,110,32,67,111,110,116,101,110,116,42,47,10,46,109,97,105,110,45,99,111,110,116,101,110,116,32,123,10,32,32,119,105,100,116,104,58,32,55,53,37,59,10,32,32,112,111,115,105,116,105,111,110,58,32,97,98,115,111,108,117,116,101,59,10,32,32,112,97,100,100,105,110,103,58,32,48,59,10,32,32,109,97,114,103,105,110,45,108,101,102,116,58,32,50,53,37,59,32,125,10,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,123,10,32,32,32,32,112,97,100,100,105,110,103,58,32,52,48,112,120,32,48,112,120,32,48,112,120,32,48,59,32,125,10,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,98,114,101,97,100,99,114,117,109,98,115,32,123,10,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,48,32,48,32,49,48,112,120,32,53,37,59,32,125,10,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,110,97,118,105,103,97,116,105,111,110,32,123,10,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,48,32,48,32,48,32,53,37,59,10,32,32,32,32,32,32,98,111,114,100,101,114,45,98,111,116,116,111,109,58,32,48,46,55,112,120,32,115,111,108,105,100,32,35,101,97,101,99,102,48,59,32,125,10,32,32,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,110,97,118,105,103,97,116,105,111,110,32,117,108,32,123,10,32,32,32,32,32,32,32,32,108,105,115,116,45,115,116,121,108,101,58,32,110,111,110,101,59,32,125,10,32,32,32,32,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,110,97,118,105,103,97,116,105,111,110,32,117,108,32,108,105,32,123,10,32,32,32,32,32,32,32,32,32,32,100,105,115,112,108,97,121,58,32,105,110,108,105,110,101,45,98,108,111,99,107,59,10,32,32,32,32,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,49,48,112,120,32,50,37,32,48,32,48,59,32,125,10,32,32,32,32,32,32,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,110,97,118,105,103,97,116,105,111,110,32,117,108,32,108,105,32,97,32,123,10,32,32,32,32,32,32,32,32,32,32,32,32,116,101,120,116,45,100,101,99,111,114,97,116,105,111,110,58,32,110,111,110,101,59,10,32,32,32,32,32,32,32,32,32,32,32,32,100,105,115,112,108,97,121,58,32,98,108,111,99,107,59,10,32,32,32,32,32,32,32,32,32,32,32,32,99,111,108,111,114,58,32,35,57,55,56,102,56,101,59,10,32,32,32,32,32,32,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,48,32,48,32,49,48,112,120,32,48,59,10,32,32,32,32,32,32,32,32,32,32,32,32,116,101,120,116,45,116,114,97,110,115,102,111,114,109,58,32,117,112,112,101,114,99,97,115,101,59,10,32,32,32,32,32,32,32,32,32,32,32,32,102,111,110,116,45,115,105,122,101,58,32,48,46,56,53,101,109,59,32,125,10,32,32,32,32,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,110,97,118,105,103,97,116,105,111,110,32,117,108,32,108,105,58,104,111,118,101,114,32,97,32,123,10,32,32,32,32,32,32,32,32,32,32,99,111,108,111,114,58,32,103,114,97,121,59,32,125,10,32,32,32,32,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,110,97,118,105,103,97,116,105,111,110,32,117,108,32,46,97,99,116,105,118,101,32,97,32,123,10,32,32,32,32,32,32,32,32,32,32,99,111,108,111,114,58,32,98,108,97,99,107,59,10,32,32,32,32,32,32,32,32,32,32,98,111,114,100,101,114,45,98,111,116,116,111,109,58,32,50,46,53,112,120,32,115,111,108,105,100,32,35,50,48,55,53,98,99,59,32,125,10,32,32,32,32,32,32,32,32,46,109,97,105,110,45,99,111,110,116,101,110,116,32,46,112,97,110,101,108,32,46,110,97,118,105,103,97,116,105,111,110,32,117,108,32,46,97,99,116,105,118,101,58,104,111,118,101,114,32,97,32,123,10,32,32,32,32,32,32,32,32,32,32,99,111,108,111,114,58,32,98,108,97,99,107,59,32,125,10,10,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,123,10,32,32,112,97,100,100,105,110,103,58,32,51,48,112,120,32,53,37,59,10,32,32,98,97,99,107,103,114,111,117,110,100,58,32,35,101,99,102,48,102,49,59,32,125,10,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,123,10,32,32,32,32,98,111,114,100,101,114,58,32,48,46,55,112,120,32,115,111,108,105,100,32,35,100,50,100,55,100,51,59,10,32,32,32,32,45,119,101,98,107,105,116,45,98,111,114,100,101,114,45,114,97,100,105,117,115,58,32,53,112,120,59,10,32,32,32,32,45,109,111,122,45,98,111,114,100,101,114,45,114,97,100,105,117,115,58,32,53,112,120,59,10,32,32,32,32,45,109,115,45,98,111,114,100,101,114,45,114,97,100,105,117,115,58,32,53,112,120,59,10,32,32,32,32,98,111,114,100,101,114,45,114,97,100,105,117,115,58,32,53,112,120,59,10,32,32,32,32,98,97,99,107,103,114,111,117,110,100,58,32,119,104,105,116,101,59,10,32,32,32,32,47,42,32,82,101,115,117,108,116,32,116,101,109,112,108,97,116,101,32,99,115,115,42,47,10,32,32,32,32,47,42,32,83,81,76,32,116,101,120,116,97,114,101,97,32,116,101,109,112,108,97,116,101,32,99,115,115,32,42,47,10,32,32,32,32,47,42,32,69,114,114,111,114,32,112,111,112,117,112,115,32,99,115,115,32,42,47,32,125,10,32,32,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,116,97,98,108,101,32,123,10,32,32,32,32,32,32,119,105,100,116,104,58,32,57,55,37,59,32,125,10,32,32,32,32,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,116,97,98,108,101,32,116,114,32,116,104,32,123,10,32,32,32,32,32,32,32,32,116,101,120,116,45,116,114,97,110,115,102,111,114,109,58,32,117,112,112,101,114,99,97,115,101,59,10,32,32,32,32,32,32,32,32,98,111,114,100,101,114,45,98,111,116,116,111,109,58,32,48,46,55,112,120,32,115,111,108,105,100,32,35,101,97,101,99,102,48,59,10,32,32,32,32,32,32,32,32,102,111,110,116,45,119,101,105,103,104,116,58,32,110,111,114,109,97,108,59,10,32,32,32,32,32,32,32,32,102,111,110,116,45,115,105,122,101,58,32,48,46,56,55,101,109,59,10,32,32,32,32,32,32,32,32,116,101,120,116,45,97,108,105,103,110,58,32,108,101,102,116,59,10,32,32,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,49,51,112,120,32,48,32,49,51,112,120,32,49,37,59,32,125,10,32,32,32,32,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,116,97,98,108,101,32,116,114,32,116,100,32,123,10,32,32,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,49,48,112,120,32,48,32,49,48,112,120,32,49,37,59,32,125,10,32,32,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,102,111,114,109,32,123,10,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,51,53,112,120,32,53,37,32,55,53,112,120,32,53,37,59,32,125,10,32,32,32,32,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,102,111,114,109,32,116,101,120,116,97,114,101,97,32,123,10,32,32,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,50,48,112,120,32,50,37,59,10,32,32,32,32,32,32,32,32,119,105,100,116,104,58,32,57,54,37,59,10,32,32,32,32,32,32,32,32,104,101,105,103,104,116,58,32,49,55,53,112,120,59,10,32,32,32,32,32,32,32,32,98,111,114,100,101,114,58,32,49,112,120,32,115,111,108,105,100,32,103,114,97,121,59,10,32,32,32,32,32,32,32,32,111,117,116,108,105,110,101,58,32,110,111,110,101,59,10,32,32,32,32,32,32,32,32,114,101,115,105,122,101,58,32,110,111,110,101,59,10,32,32,32,32,32,32,32,32,102,111,110,116,45,115,105,122,101,58,32,49,46,49,101,109,59,32,125,10,32,32,32,32,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,102,111,114,109,32,46,98,116,110,45,115,117,98,109,105,116,32,123,10,32,32,32,32,32,32,32,32,102,108,111,97,116,58,32,114,105,103,104,116,59,32,125,10,32,32,32,32,46,99,111,110,116,97,105,110,101,114,45,99,117,115,116,111,109,32,46,105,110,116,101,114,97,99,116,105,118,101,32,46,101,114,114,111,114,95,112,111,112,117,112,32,123,10,32,32,32,32,32,32,98,97,99,107,103,114,111,117,110,100,58,32,35,99,57,51,55,53,54,59,10,32,32,32,32,32,32,99,111,108,111,114,58,32,119,104,105,116,101,59,10,32,32,32,32,32,32,112,97,100,100,105,110,103,58,32,49,48,112,120,32,48,32,49,48,112,120,32,49,37,59,32,125,10,10,46,109,97,105,110,45,116,114,97,110,115,105,116,105,111,110,32,123,10,32,32,45,119,101,98,107,105,116,45,116,114,97,110,115,105,116,105,111,110,58,32,109,97,114,103,105,110,45,108,101,102,116,32,48,46,52,115,44,32,119,105,100,116,104,32,48,46,52,115,59,10,32,32,45,109,111,122,45,116,114,97,110,115,105,116,105,111,110,58,32,109,97,114,103,105,110,45,108,101,102,116,32,48,46,52,115,44,32,119,105,100,116,104,32,48,46,52,115,59,10,32,32,45,109,115,45,116,114,97,110,115,105,116,105,111,110,58,32,109,97,114,103,105,110,45,108,101,102,116,32,48,46,52,115,44,32,119,105,100,116,104,32,48,46,52,115,59,10,32,32,116,114,97,110,115,105,116,105,111,110,58,32,109,97,114,103,105,110,45,108,101,102,116,32,48,46,52,115,44,32,119,105,100,116,104,32,48,46,52,115,59,32,125,10,10,46,115,105,100,101,45,116,114,97,110,115,105,116,105,111,110,32,123,10,32,32,45,119,101,98,107,105,116,45,116,114,97,110,115,105,116,105,111,110,58,32,119,105,100,116,104,32,48,46,52,115,59,10,32,32,45,109,111,122,45,116,114,97,110,115,105,116,105,111,110,58,32,119,105,100,116,104,32,48,46,52,115,59,10,32,32,45,109,115,45,116,114,97,110,115,105,116,105,111,110,58,32,119,105,100,116,104,32,48,46,52,115,59,10,32,32,116,114,97,110,115,105,116,105,111,110,58,32,119,105,100,116,104,32,48,46,52,115,59,32,125,10,10,47,42,35,32,115,111,117,114,99,101,77,97,112,112,105,110,103,85,82,76,61,115,116,121,108,101,46,99,115,115,46,109,97,112,32,42,47,10}}
