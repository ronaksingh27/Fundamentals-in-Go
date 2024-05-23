package main

import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("If you want to be different from others , you can do same things as others ")
	fmt.Println("Switch and Case")

	diceNumber := rand.Intn(6) + 1

	switch diceNumber{
	case 1 :
		fmt.Println("Dice value is 1 and you can open ")
	case 2 : 
		fmt.Println("You can move 2 spot")
		fallthrough
	case 3 :
		fmt.Println("YOu can move 3 spot")
		fallthrough
	case 4 :
		fmt.Println("You can move 4 spot ")
	case 5 :
		fmt.Println("You can move 5 spot")
	case 6 :
		fmt.Println("You can move 6 spots , Roll the dice again")
		diceNumber = rand.Intn(6) + 1
		fmt.Println("Dice number ",diceNumber)
		for diceNumber == 6 {
			fmt.Println("You can move 6 spots , Roll the dice again")
			diceNumber = rand.Intn(6) + 1
			fmt.Println("Dice number ",diceNumber)
		}
	default :
		fmt.Println("Number is not on die")
	}

}
