package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Pointers")

	myNumber := 5;

	var ptr = &myNumber

	// (&) --> address of operator
	// (*) --> dereference operator ie gives the value stored at the address

	fmt.Println("Value of ptr is: ", ptr)
	fmt.Println("Value of ptr is: ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("Value of myNumber is: ", myNumber)

}