package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome to our pizza app ")
	fmt.Println("Please rate our pizza pp from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n');

	fmt.Println("Thanks for the rating , ",input);

	numRating , err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if( err != nil ){
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to you rating " , numRating + 1);
	}
}