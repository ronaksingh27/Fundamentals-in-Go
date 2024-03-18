package main

import "fmt"

func main(){
	fmt.Println("loops-break-continue-goto begins")

	days := []string{"sunday", "monday","tuesday","wednesday","thu","fri","sat"}

	for _,value := range days{
		fmt.Printf("%v ", value)
	}
	fmt.Println("")
	
	for num := 10 ; num < 20 ; num++ {

		if num == 13 {
			goto lco
		}
		if num == 15{
			break
		}
		fmt.Printf("%v ", num)
	}

	lco : fmt.Println("Goto executed")

	
}