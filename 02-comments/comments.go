package main

import (
	"fmt"
)

func main() {
	// Print “Hello, World!” to console
	fmt.Println("Hello, World!")

	z := x % 2 // Get the modulus of x
}

// First line of a block comment
// Second line of a block comment

/*
Everything here
will be considered
a block comment
*/

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use TrimSuffix instead.
func TrimRight(s, cutset string) string {
	if s == "" || cutset == "" {
		return s
	}
	return TrimRightFunc(s, makeCutsetFunc(cutset))
}
