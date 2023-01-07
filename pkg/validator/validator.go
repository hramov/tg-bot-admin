package validator

import (
	"fmt"
	"github.com/badoux/checkmail"
	"net"
)

func ValidateEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	return nil
}

func ValidatePassword(password string) error {
	//return passwordValidator.Validate(password, 50)
	return nil
}

func ValidateFio(fio string) error {
	return nil
}

func ValidateIp(ip string) error {
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("ip is wrong")
	}
	return nil
}
