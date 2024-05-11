package main

// golang does not have the try catch system to catch errors

// it expects to treat the problems/error to be treated like errors or true/false scenarios

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok syntax ||  error ok syntax

	input, error := reader.ReadString('\n') 

	// ----->> \n is the ascii value of new line and here using "\n" means that we want to read until the user inputs "\n"
	// input , _ := reader.ReadString('\n') ------->> means that we are not interested in the error
	// _, error := reader.ReadString('\n') ------->> means that we are not interested in the input
	// input , error := reader.ReadString('\n') ------->> means that we are interested in both input and error
	
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T ", input)
	fmt.Println("The error , ",error)

}