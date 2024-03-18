package main

import "fmt"

func main(){
	fmt.Println("structs in golang")

	user1 := User{"ronak","ronaksingh2708@gmail.com",true,21}
	fmt.Println("Details of Ronak ",user1)
	fmt.Printf("Name : %v , from %+v \n",user1.Name,user1)

	user1.GetName()
	user1.NewEmail()


}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetName(){
	fmt.Printf("Users name is %v\n ", u.Name )
}

// N -> in NewEmail make it public
func (u User) NewEmail(){
	u.Email = "go@dev"
	fmt.Printf("New Email %v \n", u.Email)
}