package mapping

import "testing"

func TestMapTo(t *testing.T) {
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
	//var userDto UserDto
	var userDto *UserDto
	err := MapTo(user, userDto)
	if err != nil {
		t.Failed()
	}

	if userDto.ID != user.ID && userDto.Name != user.Name {
		t.Failed()
	}
}
