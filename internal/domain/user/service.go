package user

import "github.com/hramov/tg-bot-admin/internal/adapters/api/user"

type Storage interface {
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) user.Service {
	return &Service{storage: storage}
}
