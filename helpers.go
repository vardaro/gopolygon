package gopolygon

/*
	This file contains convenience functions to extract the address from a literal.

	This is supposed to make it easier to use the query-as-a-struct design I've chosen.
	In Go it's pretty much impossible to do something like:
		a := &5
	Instead, you'd need to do something to the effect of
		other := 5
		a := %other
	Or do so with a function, which is what the file aims to shorten (it can get pretty verbose)
	https://willnorris.com/2014/05/go-rest-apis-and-pointers/
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
