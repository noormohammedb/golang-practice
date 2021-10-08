/*
	functions only pass copy of vairables
	for pass actual variable use pointers
*/
package main

import "fmt"

func main() {
	user01 := User{"firstUser", true}

	user01.GetStatus()
	user01.ChangeStatus(false)
	user01.GetStatus()

}

type User struct {
	name   string
	status bool
}

func (newUser User) GetStatus() {
	fmt.Println(newUser.status)
}

func (newUser User) ChangeStatus(newStatus bool) {
	newUser.status = newStatus
	fmt.Println("newUser : ", newUser)
}
