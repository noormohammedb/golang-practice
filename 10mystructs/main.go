package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	user01 := User{"firstUser", "user@email.com", "blablapass", true, 33}
	fmt.Println(user01)
	fmt.Printf("%+v\n", user01)
	fmt.Printf("name: %v", user01.name)
}

type User struct {
	name     string
	email    string
	password string
	status   bool
	age      int
}
