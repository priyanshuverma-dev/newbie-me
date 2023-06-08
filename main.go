// package name
package main

/*
**	Importing libs
**	import "fmt"
**	OR
**	import ("fmt")
 */

import (
	`fmt` // lib to get printing stuff
)

//declaring function with `func` and then `functionName`

func main() {
	// Printing `Hello, Go!` in terminal
	// To use fmt.Println("")
	fmt.Println("Hello, Go!")

	// variables

	var wholeNumber int = 9          // Integer
	var Name string = "Priyanshu"    // String
	var DecimalNo float32 = 9.9      // Decimal Number of 32 bytes mem
	var DecimalNo1 float64 = 9.99999 // Decimal Number of 64 bytes mem
	var isProgrammer bool            // boolean True/False

	// Printing All Variables With Printf()
	// `%v` variable
	// `%T` type of varible
	fmt.Printf("%v,%v,%v,%v,%v", wholeNumber, Name, DecimalNo, DecimalNo1, isProgrammer)

}
