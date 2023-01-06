package authTest

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/src/domain/auth"
)

type AuthRegistryMock struct{}

func NewAuthRegistryMock() AuthRegistryMock {
	return AuthRegistryMock{}
}

var admin = auth.Login{
	Id:       120,
	Email:    "admin@admin.ru",
	Password: "$2a$10$WeYBd9gUitJa1QiNSh/I8OQzCm/uL0.AdHGG1VnlkX037/Q4Xj5Ri",
}

func (ar AuthRegistryMock) GetCandidate(ctx context.Context, email string) (*auth.Login, error) {
	if email == "admin@admin.com" {
		return nil, fmt.Errorf("some database error")
	}

	if email == admin.Email {
		return &admin, nil
	}

	return &auth.Login{
		Id: 0,
	}, nil
}
