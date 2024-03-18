package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating fo our pizza : ")

	// comma ok || err ok
	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating , \n",input)
	fmt.Println("Type of rating is %T \n",input)


}