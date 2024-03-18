package main

import "fmt"

func main(){

	defer fmt.Println("Two")
	defer fmt.Println("One")

	fmt.Println("Hello")

	defer myDefer()

}

func myDefer(){
	for i := 0 ; i < 10 ; i++ {
		fmt.Printf("%v ", i)
	}
}