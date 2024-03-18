package main

import "fmt"

func main(){

	fmt.Println("Fuctions , Lets go Young Man , Work Hard")
	add := adder(5,7)
	fmt.Printf("add : %v \n", add)

	add , message , boolVal := addMany(1,2,3,4,5)
	fmt.Printf("add : %v , message : %v , bool : %v\n", add , message,boolVal)

}


func adder(valOne int  , valTwo int ) int {
	return valOne + valTwo;
}

func addMany( values ...int) (int , string , bool) {
	
	total := 0
	for _ , value := range values{
		total += value
	}

	return total , "Hello from addMany", true
}