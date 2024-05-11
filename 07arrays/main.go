package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	// in golang we have initialize the array with some default length/size of the 
	// array while declaration . This is mandatory according to the 
	// language specifications

	var fruitList [4]string 

	fruitList[0] = "Apple";
	fruitList[1] = "Orange";
	fruitList[3] = "Peach";

	fmt.Println("Fruit list is: ", fruitList)

	for i := 0; i < len(fruitList); i++ {
		fmt.Println("Element at index", i, "is", fruitList[i])	
	}

	var vegetableList = [5]string {"Potato", "Beans", "Cabbage", "Carrot", "Onion"}

	fmt.Println("Vegetable list is: ", vegetableList)

}