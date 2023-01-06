package jwt

import (
	"testing"
)

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

//func TestCreateTokenPair(t *testing.T) {
//	//id, _ := uuid.Parse("4b10aaf4-7c32-4bef-8fc3-fd7e00ef8bc7")
//
//	user := ""
//	_, _, err := CreateToken(user)
//
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}
//
//func TestVerifyAccessToken(t *testing.T) {
//	//_, _ := uuid.Parse("4b10aaf4-7c32-4bef-8fc3-fd7e00ef8bc7")
//	user := ""
//	at, _, err := CreateToken(user)
//
//	_, err = VerifyToken(at, 0)
//
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}

//func TestVerifyRefreshToken(t *testing.T) {
//	//id, _ := uuid.Parse("4b10aaf4-7c32-4bef-8fc3-fd7e00ef8bc7")
//	user := ""
//	_, rt, err := CreateToken(user)
//
//	_, err = VerifyToken(rt, 0)
//
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}

func TestCreateAuthCode(t *testing.T) {
	code := CreateAuthCode()
	t.Log(code)
}

func TestCreateClientId(t *testing.T) {
	code := CreateClientId()
	code2 := CreateClientId()

	if code == code2 {
		t.Errorf("equal codes")
	}
}

func TestCreateClientSecret(t *testing.T) {
	code, err := CreateClientSecret()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(code)
}
