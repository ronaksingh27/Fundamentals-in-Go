package main

import "fmt"

var jwtToken  = 213111
const LoginToken = "aasdfgh";// L -> Makes LogiToken public
func main(){

	var username string = "hitesh";
	// fmt.Println("variables");
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n",username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n",smallVal)
	

	var floatVal float32 = 255.123456789;
	fmt.Println(floatVal)
	fmt.Printf("variable is of type : %T \n",floatVal)

	//default and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("variable is of type : %T \n",anotherVar)

	//implicit type
	var website = "ronaksingh.com"
	fmt.Println(website)

	//walrus operator / no var style
	address := "VNIT";
	fmt.Printf("variable is of type : %T \n",address)
	
}