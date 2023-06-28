package main

import "fmt"

type user struct {
	firstName string
	LastName  string
	Age       int
}

func (user user) GetFullName() string {
	return user.firstName + " " + user.LastName
}

func (user *user) SetFirstName(firstName string) {
	user.firstName = firstName
	fmt.Printf("copy user: %+v\n", user)
}

func main() {
	// user1 := user{
	// 	firstName: "Somkiat",
	// 	LastName:  "Chawkamud",
	// 	Age:       30,
	// }
	// user2 := user{
	// 	firstName: "Dodo",
	// 	LastName:  "Chawkamud",
	// 	Age:       30,
	// }

	// users := []user{user1, user2}

	// for _, v := range users {
	// 	println(v.GetFullName())
	// }
	// user.firstName = "ittikorn"
	// user.Age = 20
	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%+v\n", user.GetFullName())

	methodPointer()
}

func methodPointer() {
	user := user{
		firstName: "Somkiat",
		LastName:  "Chawkamud",
		Age:       30,
	}

	user.SetFirstName("ittikorn")

	fmt.Printf("%+v\n", user)
}
