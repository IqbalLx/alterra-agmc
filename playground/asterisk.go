package main

import (
	"fmt"
)

func main() {
	var a = 5
	var p = &a // p holds variable a's memory address
	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)

	// Let's change a value (using the initial variable or the pointer)
	// *p = 3 // using pointer
	// a = 3 // using initial var

	// fmt.Printf("Address of var a: %p\n", p)
	// fmt.Printf("Value of var a: %v\n", *p)

	var c = &a
	fmt.Printf("Address of var a: %p\n", c)
}
