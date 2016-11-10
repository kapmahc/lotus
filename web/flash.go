package web

//Notice notice flash message
func Notice(val interface{}) H {
	return H{"notice": val}
}

//Error error flash message
func Error(val interface{}) H {
	return H{"error": val}
}
