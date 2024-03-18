package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url  = "https://lco.dev"
func main(){
	fmt.Println("Web Requests : HTTP")

	response , err := http.Get(url)
	if( err != nil){
		panic(err)
	}


	fmt.Printf("Response is of type : %T \n", response)
	defer response.Body.Close()//callers repsonsibilty to close connection

	databyte , err := ioutil.ReadAll(response.Body)
	if( err != nil){
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}