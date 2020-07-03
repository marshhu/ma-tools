package js

import (
	"fmt"
	"testing"
)

func Test_PasswordEnc(t *testing.T) {
	salt := "1379BFX"
	password := "password"
	except := salt + password + salt
	encPwd := PasswordEnc(salt, password)
	fmt.Println(encPwd)
	if encPwd != except {
		t.FailNow()
	}
}
