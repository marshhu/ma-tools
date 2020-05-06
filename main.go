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

	type AddressB struct {
		Country string
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
		Address AddressB
	}

	user := User{1, "小明", 1, "18800188001", Address{"中国", "深圳"}}
	fmt.Printf("user:%v\n", user)
	userDto := &UserDto{}
	err := mapping.MapTo(user, userDto)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("userDto:%v\n", userDto)

	users := []User{
		{1, "小明", 1, "18800188001", Address{"中国", "深圳"}},
		{2, "小红", 0, "18800188002", Address{"中国", "广州"}},
		{3, "小李", 0, "18800188003", Address{"中国", "武汉"}},
		{4, "小张", 1, "18800188004", Address{"中国", "北京"}},
	}
	fmt.Printf("users:%v\n", users)
	var userDtos []UserDto
	err = mapping.MapTo(users, &userDtos)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("userDtos:%v\n", userDtos)
}
