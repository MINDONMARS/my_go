package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	if ( a && b ) {
		fmt.Printf("1 true\n")
	} else {
		fmt.Printf("1 false\n")
	}
	if ( a || b ) {
		fmt.Printf("2 true\n")
	} else {
		fmt.Printf("2, false\n")
	}
	if ( !(a && b) ) {
		fmt.Printf("3 true\n")
	} else {
		fmt.Printf("3 false")
	}
}
