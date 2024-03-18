package main

import "fmt"

func main(){
	fmt.Println("Welcome to maps")

	languages := make( map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	delete(languages,"JS")

	fmt.Println("maps : ",languages)
	fmt.Println("maps : ",languages["PY"])

	for _,value := range(languages){
		fmt.Println("For key v , value is ",value)
	}
}