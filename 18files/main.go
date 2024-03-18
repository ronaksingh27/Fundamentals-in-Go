package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/* We can use String class getBytes() method to encode the string into a sequence of bytes using the platform’s default charset.*/


func main(){

	if( len(os.Args) != 2 ){
		fmt.Println("Please provide a filename")
		return
	}

	content := "Hello this is Ronak , leanring to write code"

	// fmt.Println(os.Args[1])
	filename := os.Args[1]
	file,err := os.Create(filename)
	checkNilErr(err)
	
	
	length, err := fmt.Fprintf( file ,content) // fileptr , text
	fmt.Println("lenght of content is : ",length)
	checkNilErr(err)

	defer file.Close()

	fmt.Fprintf( file , " Writing to file ") // fileptr , text , filename

	readFile(filename)
}


func checkNilErr(err error){
	if( err != nil ){
		panic(err)
	}
}

func readFile(filename string){

	// fmt.Println(filename)
	databyte , err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text inside the file is ",string(databyte))
}
