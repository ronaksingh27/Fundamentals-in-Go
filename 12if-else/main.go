package main

import "fmt"

func main(){
	fmt.Println("Lets play with if-else")

	num := 27

	if num < 27{
		fmt.Printf("Number is less than %v \n",num)
	}else if num == 27{
		fmt.Printf("Number is equal to %v \n", num)
	}else{
		fmt.Printf("Number is greater than %v \n", num)
	}

	if num % 2 == 0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	if num := 7 ; num == 7{
		fmt.Printf("Ronak 2%v08 \n",num )
	}

	// fmt.Printf("Ronak 2%v08 \n",num )


}