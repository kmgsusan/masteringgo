package main

import (
	"io"
	"os"
)

// dd
func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = " Pleasse "
	} else {
		// myString = arguments[1]
		for i, x := range arguments {
			if i > 0 {
				myString += x + ", "
			}
		}
	}

	io.WriteString(os.Stdout, myString)
}
