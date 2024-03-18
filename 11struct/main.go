package main

import "fmt"

func main(){
	fmt.Println("structs in golang")

	ronak := User{"ronak","ronaksingh2708@gmail.com",true,21}
	fmt.Println("Details of Ronak ",ronak);
	fmt.Printf("Name : %v , from %+v ",ronak.Name,ronak);



}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}