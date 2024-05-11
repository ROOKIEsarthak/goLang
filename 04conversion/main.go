package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out Pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)


	numRating,err := strconv.ParseFloat(strings.Split(input, "\n")[0], 64)

	if(err != nil){
		fmt.Println("Error",err)
	} else {
		fmt.Println("Added 1 to your rating: ",numRating+1)
	}
}