package gopolygon

/*
	This file contains convenience functions to extract the address from a literal.

	This is supposed to make it easier to use the query-as-a-struct design I've chosen
*/

// String just returns the address of the string
// Meant to be called address := gopolygon.String("my literal string")
func String(str string) *string {
	return &str
}

// Bool returns the address of the input
func Bool(b bool) *bool {
	return &b
}
