package main

import (
	"fmt"
	"ma-tools/mapping"
)

func main() {
	type User struct {
		ID     int
		Name   string
		Gender int
		Tel    string
	}

	type UserDto struct {
		ID     int
		Name   string
		Avatar string
	}

	user := User{1, "小明", 1, "18800188001"}
	//userDto := &UserDto{}
	var userDto *UserDto
	err := mapping.MapTo(user, userDto)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("userDto:%v", userDto)
}
