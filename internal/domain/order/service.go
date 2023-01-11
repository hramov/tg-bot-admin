package order

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/config"
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type Storage interface {
}

type Service interface {
	GetAll(ctx context.Context) ([]*Order, appError.IAppError)
	GetBy(ctx context.Context, field, value string) ([]*Order, appError.IAppError)
	GetUserOrders(ctx context.Context, userId int) ([]*InfoForUser, appError.IAppError)
	Create(ctx context.Context, dto ShallowWeedProduct) (*int, appError.IAppError)
	ChangeStatus(ctx context.Context, statusId int) (*int, appError.IAppError)
}

type service struct {
	validator *validator.Validate
	storage   Storage
	logger    *logging.Logger
	cfg       *config.Config
}

func NewService(storage Storage, validator *validator.Validate, logger *logging.Logger, cfg *config.Config) Service {
	return &service{storage: storage, validator: validator, logger: logger, cfg: cfg}
}

func (s *service) GetAll(ctx context.Context) ([]*Order, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetBy(ctx context.Context, field, value string) ([]*Order, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetUserOrders(ctx context.Context, userId int) ([]*InfoForUser, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Create(ctx context.Context, dto ShallowWeedProduct) (*int, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *service) ChangeStatus(ctx context.Context, statusId int) (*int, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}
