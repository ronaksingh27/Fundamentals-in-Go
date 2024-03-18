package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myUrl = "http://localhost:2708/post"
func main(){
	fmt.Println("Welcome  to web verb video")
	
	// PerformGetRequest(myUrl)
	// PerformPostJsonRequest(myUrl)
	PerfomrPostFormRequest()
}

func PerformGetRequest( myUrl string ){
	response , err := http.Get(myUrl)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.ContentLength)

	var responseString strings.Builder
	content , err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	bytecount , err := responseString.Write(content)
	if err != nil{
		panic(err)
	}

	fmt.Println("bytecount is : ",bytecount)
	fmt.Println(responseString.String())

	// content := string(databyte)
	// fmt.Println(content)
}

func PerformPostJsonRequest( myUrl string ){
	requestBody := strings.NewReader(`
		{
			"coursename" : "Lets learn go lang ",
			"price" : 0,
			"platform" : "learncodeonline"
		}
	`)

	response , err := http.Post(myUrl,"application/json",requestBody)
	if err!= nil{
		panic(err)
	}
	defer response.Body.Close()

	content ,err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(content))
}

func PerfomrPostFormRequest(){
	const myUrl = "http://localhost:2708/postform"

	//format data
	data := url.Values{}
	data.Add("JS","Javascript")
	data.Add("RB","Ruby")
	data.Add("GO","GOlang")

	response , err := http.PostForm(myUrl,data)
	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	content , err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))



}