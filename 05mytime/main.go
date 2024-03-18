package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime);
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2024, time.August,27,12,10,27,17,time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006"))
}