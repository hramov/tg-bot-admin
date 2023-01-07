package user

import (
	"context"
)

type Storage interface {
	GetById(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Get(ctx context.Context, limit, offset int) ([]*User, error)
}

type Service struct {
	storage Storage
}

type IService interface {
	GetAll(ctx context.Context, limit, offset int) ([]*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetById(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, dto *CreateDto) (int, error)
	Update(ctx context.Context, dto *UpdateDto) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

func NewService(storage Storage) IService {
	return &Service{storage: storage}
}

func (s Service) GetAll(ctx context.Context, limit, offset int) ([]*User, error) {
	return s.storage.Get(ctx, limit, offset)
}

func (s Service) GetByEmail(ctx context.Context, email string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetById(ctx context.Context, id int) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create(ctx context.Context, dto *CreateDto) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(ctx context.Context, dto *UpdateDto) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, id int) (int, error) {
	//TODO implement me
	panic("implement me")
}
