package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{fullName}
}

var AllNames []string

func GetOrAdd(s string) uint8 {
	for i := range AllNames {
		if AllNames[i] == s {
			return uint8(i)
		}
	}
	AllNames = append(AllNames, s)
	return uint8(len(AllNames) - 1)
}

type User2 struct {
	names []uint8
}

func NewUser2(fullname string) *User2 {

	result := User2{}
	parts := strings.Split(fullname, " ")
	for _, p := range parts {
		result.names = append(result.names, GetOrAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, AllNames[id])
	}
	return strings.Join(parts, " ")
}

func main() {

	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	alsejane := NewUser("Jane Smith")
	fmt.Println(john, jane, alsejane)

	john1 := NewUser2("John Doe")
	jane1 := NewUser2("Jane Doe")
	alsejane1 := NewUser2("Jane Smith")

	fmt.Println(john1, jane1, alsejane1)
	fmt.Println(john1.FullName())

}
