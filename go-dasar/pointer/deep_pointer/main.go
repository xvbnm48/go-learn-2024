package main

import "fmt"

// when
// 1. need to update struct field using pointer
// 2. when we want to optimaze memory usage, using pointer because pointer = 8 byte, struct = 24 byte
type User struct {
	Email string
	Name  string
	Age   int
}

func getUser() (*User, error) {
	return nil, fmt.Errorf("error")
}

//	func getUser() (User, error) {
//		return User{}, fmt.Errorf("error")
//	}
func (u User) email() string {
	return u.Email
}

// version 1 using method func
func (u *User) UpdateEmail(email string) {
	u.Email = email
}

// version 2
func UpdateEmail2(u *User, email string) {
	u.Email = email
}

func main() {
	user := User{
		Email: "vini@hd.com",
		Name:  "vini",
		Age:   21,
	}
	user.UpdateEmail("vini@hd.co.id")
	fmt.Println(user.email())
}
