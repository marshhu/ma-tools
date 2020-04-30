package main

import (
	"fmt"
	"ma-tools/mapping"
)

func main() {
	type Address struct {
		Country string
		City    string
	}
	type User struct {
		ID      int
		Name    string
		Gender  int
		Tel     string
		Address Address
	}

	type UserDto struct {
		ID      int
		Name    string
		Avatar  string
		Address Address
	}

	user := User{1, "小明", 1, "18800188001", Address{"中国", "深圳"}}
	userDto := &UserDto{}
	err := mapping.MapTo(user, userDto)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("userDto:%v", userDto)
}
