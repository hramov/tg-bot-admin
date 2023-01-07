package utils

import "testing"

func TestIsEqualClientIp_Localhost(t *testing.T) {
	serviceIp := "http://localhost:4001"
	clientIp := "127.0.0.1"

	equal, err := IsEqualClientIp(serviceIp, clientIp)
	if err != nil {
		t.Error(err)
	}

	if !equal {
		t.Error("should be equal")
	}
}

func TestIsEqualClientIp_Remote(t *testing.T) {
	serviceIp := "http://62.113.111.13:4001"
	clientIp := "62.113.111.13"
	equal, err := IsEqualClientIp(serviceIp, clientIp)
	if err != nil {
		t.Error(err)
	}
	if !equal {
		t.Error("should be equal")
	}
}
