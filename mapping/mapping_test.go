package mapping

import (
	"testing"
	"time"
)

type Address struct {
	Country string
	City    string
}

type AddressB struct {
	Country string
}

type Hobby struct {
	Name  string
	Level int
}

type HobbyB struct {
	Name string
}

type Model struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Model
	Name    string
	Gender  int
	Tel     string
	Address Address
	Hobbies []Hobby
}

type UserDto struct {
	Id        int
	Name      string `ignore:"true"` //映射时，忽略此字段
	Avatar    string
	Phone     string `mappingField:"Tel"` //指定映射字段名
	Address   AddressB
	Hobbies   []HobbyB
	CreatedAt time.Time
	UpdatedAt string `timeFormat:"2006-01-02"` //指定映射时，时间格式
}

func GetData() []User {
	var users []User
	user1 := User{Model{1, time.Now(), time.Now()}, "小明", 1, "18800188001",
		Address{"中国", "深圳"},
		[]Hobby{{Name: "游泳", Level: 2}, {Name: "篮球", Level: 4}},
	}
	user2 := User{Model{2, time.Now(), time.Now()}, "小红", 0, "18800188002",
		Address{"中国", "深圳"},
		[]Hobby{{Name: "游泳", Level: 2}, {Name: "篮球", Level: 4}},
	}
	user3 := User{Model{3, time.Now(), time.Now()}, "小李", 0, "18800188003",
		Address{"中国", "深圳"},
		[]Hobby{{Name: "游泳", Level: 2}, {Name: "篮球", Level: 4}},
	}
	user4 := User{Model{4, time.Now(), time.Now()}, "小张", 1, "18800188004",
		Address{"中国", "深圳"},
		[]Hobby{{Name: "游泳", Level: 2}, {Name: "篮球", Level: 4}},
	}
	users = append(users, user1, user2, user3, user4)
	return users
}

func CheckResult(user *User, userDto *UserDto) bool {
	if userDto.Id != user.ID {
		return false
	}
	//if userDto.Name != user.Name {
	//	return false
	//}
	if userDto.Phone != user.Tel {
		return false
	}
	if userDto.Address.Country != user.Address.Country {
		return false
	}

	if userDto.CreatedAt.Format("2006-01-02 15:04:05") != user.CreatedAt.Format("2006-01-02 15:04:05") || userDto.UpdatedAt != user.UpdatedAt.Format("2006-01-02") {
		return false
	}
	if len(userDto.Hobbies) != len(user.Hobbies) {
		return false
	}
	for i, hobby := range userDto.Hobbies {
		if hobby.Name != user.Hobbies[i].Name {
			return false
		}
	}
	return true
}

func TestMapToStruct(t *testing.T) {
	data := GetData()
	user := data[0]
	userDto := &UserDto{}
	err := MapTo(user, userDto)
	if err != nil {
		t.FailNow()
	}
	if !CheckResult(&user, userDto) {
		t.FailNow()
	}
}

func TestMapToSlice(t *testing.T) {
	data := GetData()
	var userDtos []UserDto
	err := MapTo(data, &userDtos)
	if err != nil {
		t.FailNow()
	}
	if len(userDtos) != len(data) {
		t.FailNow()
	}

	for i := 0; i < len(data); i++ {
		userDto := userDtos[i]
		user := data[i]
		if !CheckResult(&user, &userDto) {
			t.FailNow()
		}
	}
}

func TestMapToSlice_Ptr(t *testing.T) {
	data := GetData()
	var userDtos []*UserDto
	err := MapTo(data, &userDtos)
	if err != nil {
		t.FailNow()
	}
	if len(userDtos) != len(data) {
		t.FailNow()
	}

	for i := 0; i < len(data); i++ {
		userDto := userDtos[i]
		user := data[i]
		if !CheckResult(&user, userDto) {
			t.FailNow()
		}
	}
}

func TestMapToArray(t *testing.T) {
	data := GetData()
	var users [4]User
	users[0] = data[0]
	users[1] = data[1]
	users[2] = data[2]
	users[3] = data[3]

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
		if !CheckResult(&user, &userDto) {
			t.FailNow()
		}
	}
}

func TestMapToMap(t *testing.T) {
	data := GetData()
	users := make(map[string]User)
	users["小明"] = data[0]
	users["小红"] = data[1]
	users["小李"] = data[2]
	users["小张"] = data[3]

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
		if !CheckResult(&user, &userDto) {
			t.FailNow()
		}
	}
}
