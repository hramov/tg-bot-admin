package crypto

import "testing"

func TestPasswordHash(t *testing.T) {
	hashed, err := CreateHashedPassword("admin")
	t.Log(hashed)
	if err != nil {
		t.Errorf("CreateHashPasswordError: %s", err.Error())
	}
}

func TestCheckPassword_Success(t *testing.T) {
	plain := "admin"
	hash, _ := CreateHashedPassword(plain)
	equal := CheckPassword(plain, hash)
	if equal == false {
		t.Errorf("Wrong result: should be true")
	}
}

func TestCheckPassword_Error(t *testing.T) {
	plain := "admin"
	hash, _ := CreateHashedPassword("plain")
	equal := CheckPassword(plain, hash)
	if equal == true {
		t.Errorf("Wrong result: should be false")
	}
}
