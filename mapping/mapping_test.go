package mapping

import (
	"testing"
)

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

func TestMapToStruct(t *testing.T) {
	user := User{1, "小明", 1, "18800188001", Address{"中国", "深圳"}}
	userDto := &UserDto{}
	err := MapTo(user, userDto)
	if err != nil {
		t.FailNow()
	}
	if userDto.ID != user.ID && userDto.Name != user.Name && userDto.Address.Country != user.Address.Country {
		t.FailNow()
	}
}

func TestMapToSlice(t *testing.T) {
	users := []User{
		{1, "小明", 1, "18800188001", Address{"中国", "深圳"}},
		{2, "小红", 0, "18800188002", Address{"中国", "广州"}},
		{3, "小李", 0, "18800188003", Address{"中国", "武汉"}},
		{4, "小张", 1, "18800188004", Address{"中国", "北京"}},
	}
	var userDtos []UserDto
	err := MapTo(users, &userDtos)
	if err != nil {
		t.FailNow()
	}
	if len(userDtos) != len(users) {
		t.FailNow()
	}

	for i := 0; i < len(users); i++ {
		userDto := userDtos[i]
		user := users[i]
		if userDto.ID != user.ID && userDto.Name != user.Name && userDto.Address.Country != user.Address.Country {
			t.FailNow()
		}
	}
}

func TestMapToArray(t *testing.T) {
	users := [4]User{
		{1, "小明", 1, "18800188001", Address{"中国", "深圳"}},
		{2, "小红", 0, "18800188002", Address{"中国", "广州"}},
		{3, "小李", 0, "18800188003", Address{"中国", "武汉"}},
		{4, "小张", 1, "18800188004", Address{"中国", "北京"}},
	}
	var userDtos [4]UserDto
	err := MapTo(users, &userDtos)
	if err != nil {
		t.FailNow()
	}
	if len(userDtos) != len(users) {
		t.FailNow()
	}

	for i := 0; i < len(users); i++ {
		userDto := userDtos[i]
		user := users[i]
		if userDto.ID != user.ID && userDto.Name != user.Name && userDto.Address.Country != user.Address.Country {
			t.FailNow()
		}
	}
}

func TestMapToMap(t *testing.T) {
	users := make(map[string]User)
	users["小明"] = User{1, "小明", 1, "18800188001", Address{"中国", "深圳"}}
	users["小红"] = User{2, "小红", 0, "18800188002", Address{"中国", "广州"}}
	users["小李"] = User{3, "小李", 0, "18800188003", Address{"中国", "武汉"}}
	users["小张"] = User{4, "小张", 1, "18800188004", Address{"中国", "北京"}}

	userDtos := make(map[string]UserDto)
	err := MapTo(users, &userDtos)
	if err != nil {
		t.FailNow()
	}
	if len(userDtos) != len(users) {
		t.FailNow()
	}

	for key := range users {
		userDto := userDtos[key]
		user := users[key]
		if userDto.ID != user.ID && userDto.Name != user.Name && userDto.Address.Country != user.Address.Country {
			t.FailNow()
		}
	}
}
