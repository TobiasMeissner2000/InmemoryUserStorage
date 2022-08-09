package user

import "fmt"

type User struct {
	Id        int //`json:"id"`
	FirstName string
	LastName  string
	Age       int
}

var (
	lastUserID int = 1000 //first user has the id 1001
)

func (u *User) Set(firstName, lastName string, age int) {
	u.SetId()
	u.FirstName = firstName
	u.LastName = lastName
	u.Age = age
}

func (u *User) SetId() {
	lastUserID++
	u.Id = lastUserID
}

func (u *User) Print() {
	fmt.Printf("Name: %s Lastname: %s Age: %d Id: %d \n", u.FirstName, u.LastName, u.Age, u.Id)
}
