package user

import "fmt"

type User struct {
	ID        int //`json:"id"`
	FirstName string
	LastName  string
	Age       int
}

func (u *User) Set(firstName, lastName string, age int) {
	u.ID = -1
	u.FirstName = firstName
	u.LastName = lastName
	u.Age = age
}

func (u *User) Print() {
	fmt.Printf("Name: %s Lastname: %s Age: %d Id: %d \n", u.FirstName, u.LastName, u.Age, u.ID)
}
