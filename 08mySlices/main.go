package main

import (
	"fmt"
	"sort"
)

// ----->> in goLang slices are used instead of arrays
// ----->> slices are reference types

func main() {

	fmt.Println("Welcome to slices in golang")

	// if you give a certain number in square brackets while declaration it becomes an array
	// if you don't give a number in square brackets while declaration it becomes a slice

	var fruitList = []string{"Apple","Tomato","Peach"}


	fmt.Printf("Fruit list is of type %T : ", fruitList)
	fmt.Println("Fruit list is: ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	// ----->> append method takes the list (fruitList) and adds values like "Mango","Banana" to the slice
	fmt.Println("Fruit list is: ", fruitList)

	fruitList = append(fruitList[:])
	fmt.Println("Fruit list is: ", fruitList)

	// new method of initialization
	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 543

	highScore = append(highScore, 555, 666, 321)
	 
	fmt.Println("High score is: ", highScore)
	sort.Ints(highScore)
	fmt.Println("Sorted High score is: ", highScore)

	fmt.Println(sort.IntsAreSorted(highScore))


	// how to remove a value from slices based on index
	var course = []string{"Java","Python","Javascript","Golang","React"}
	fmt.Println("courses are : ", course)

	var index int = 2;
	course = append(course[:index],course[index+1:]...)
	fmt.Println("courses are : ", course)


}


