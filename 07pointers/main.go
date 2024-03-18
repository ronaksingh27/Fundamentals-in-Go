package main

import "fmt"

func main(){
	fmt.Println("Welcome to pointers");

	myNum := 27
	var ptr  = &myNum

	fmt.Println("value of pointer is ",*ptr)
}