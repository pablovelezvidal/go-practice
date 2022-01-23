package main

import (
	"fmt"
)

type User struct {
	Id      int
	Name    string
	Age     int
	Friends []User
}

func (u User) SayHello() {
	fmt.Printf("Hello me llamo %s \n", u.Name)
}

func (u *User) AddFriends(friend User) {
	u.Friends = append(u.Friends, friend)
}

func (u User) findFriend(id int) int {
	for i, f := range u.Friends {
		if f.Id == id {
			return i
		}
	}
	return -1
}

func (u *User) RemoveFriends(id int) {
	index := u.findFriend(id)
	if index < 0 {
		return
	}

	u.Friends = append(u.Friends[:index], u.Friends[index+1:]...)
}

func (u *User) ListFriends() {
	for i, f := range u.Friends {
		fmt.Printf("The friend %d is: %s \n", i, f.Name)
	}
}

func main() {
	martha := User{Id: 1, Name: "martha", Age: 12}
	pedro := User{Id: 2, Name: "pedro", Age: 15}
	juan := User{Id: 3, Name: "juan", Age: 16}

	martha.SayHello()
	martha.AddFriends(pedro)
	martha.AddFriends(juan)

	martha.ListFriends()

	martha.RemoveFriends(pedro.Id)

	martha.ListFriends()
}
