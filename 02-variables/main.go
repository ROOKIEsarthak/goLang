package main
import "fmt"

//  ------->> in golang there is no need to declare variable type
// ---------->> in golang we declare variables with small letters
// --------->> in golang we use var keyword to declare variable
// ---------->>> in golang we use :=(wallarus) operator to declare variable


// ------>>> in golang declaring a variable with first letter capital 
// marks it as a public variable which is accessible to other packages .
const LoginToken string = "adjnakjxiasndknlkaaxslkmad"

func main() {
	var username string = "Sarthak"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)


	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)


	var smallFloat float64 = 255.455430129374928749823409283409324090329092384098
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)


	// default values and some aliases

	//-----> go puts default value as 0 in int,float and null in string


	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "https://www.google.com"
	fmt.Println(website)

	// no var style

	noOfUser := 300000
	fmt.Println(noOfUser)



}