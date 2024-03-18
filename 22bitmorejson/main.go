package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string `json:"coursename"`
	Price string
	Platform string `json:"website"`
	Password string	`json:"-"`
	Tags []string	`json:"tags,omitempty"`
}
func main(){
	fmt.Println("Welcome to creaating JSON form data")
	EncodeJson()
	// DecodeJson()
	
}

func EncodeJson(){
	lcoCourses := []course{
		{"Javascript","199","js.dev","js123",[]string{"js","easy","conceptual"} },
		{"Ruby","299","js.dev","rb123",[]string{"ruby","easy","conceptual"} },
		{"GO","399","go.dev","go123",[]string{"golang","easy","conceptual"} },
		{"C++","499","c++.dev","c++123",[]string{"C++","easy","conceptual"} },
		{"C","199","c.dev","c123",[]string{"C","easy","conceptual"} },
	}


	finalJson , err := json.MarshalIndent(lcoCourses,"","\t")
	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"coursename": "GO",
		"Price": "399",
		"website": "go.dev",
		"tags": [
				"golang",
				"easy",
				"conceptual"
		]
	}
	`)
	//Need to verify if above was in json format

	// var lcoCourse course

	// checkValid := json.Valid(jsonDataFromWeb)

	// if checkValid{
	// 	fmt.Println("JSON was valid")
	// 	json.Unmarshal(jsonDataFromWeb,&lcoCourse)
	// 	fmt.Printf("%#v\n",lcoCourse)
	// }else{
	// 	fmt.Println("JSON was not valid")
	// }


	//some cases in which you jsut add data to key value pairs
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key , value := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is : %T \n", key,value,value)
	}

}