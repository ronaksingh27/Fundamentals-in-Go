package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://learncodeonline:2708/learn?coursename=reactjs&payment=online"
func main(){
	fmt.Println("Welcome to handle urls in golang")
	fmt.Println(myUrl)

	//parsing the url
	result , err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.Hostname())
	fmt.Println(result.RawQuery)

	//stores the paramtertes as key , value pairs
	qparams := result.Query()
	fmt.Println(qparams["coursename"])

	for _,value := range(qparams){
		fmt.Printf("Params is %v \n",value)
	}

	//Constructing a URL from Scratch
	partsOfUrl := &url.URL{
		Scheme : "https",
		Host : "lco.dev",
		Path : "tutcss",
		RawPath: "user-ronak",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
