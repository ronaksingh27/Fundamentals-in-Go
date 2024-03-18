package main

import (
	"fmt"
	"slices"
)

func main(){

	/******* SLICES *********/
	fmt.Println("Welcome to slcies")
	var fruitList []string
	fmt.Printf("FruitList is of type : %T \n",fruitList);

	fruitList = append(fruitList , "Mango" , "banana","apple","grapes","Peach")
	fmt.Println("FruitList is ",fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("FruitList is ",fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("FruitList is ",fruitList)


	highScores := make( []int, 5)
	highScores[0] = 0
	highScores[1] = 10
	highScores[2] = 20
	highScores[3] = 30
	highScores[4] = 40

	fmt.Println("HighScores : ",highScores)

	highScores = append(highScores, 50,60,70)
	highScores = append(highScores, 50,60,70)
	fmt.Println("HighScores : ",highScores)


	slices.Sort(highScores)
	fmt.Println("HighScores : ",highScores)

	var courses []string
	courses = append( courses, "javascript" , "C++","ruby on rails ", "golang","java")
	
	var index int = 2;
	courses = append( courses[:index] , courses[index+1:]...)
	fmt.Println("courses : ",courses)

	/******** ARRAYS ***********/
	// fmt.Println("Welcome to array")

	// var fruitList [5]string

	// fruitList[0] = "Apple"
	// fruitList[1] = "banana"
	// fruitList[4] = "grapes"

	// fmt.Println("fruitslist is : ",fruitList)
	// fmt.Println("fruitslist is : ",len(fruitList))

	// var vegList [5]string
	// fmt.Println("Veg list ",len(vegList))
}