package order

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/config"
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type IStorage interface {
}

type IService interface {
	GetAll(ctx context.Context) ([]*Order, appError.IAppError)
	GetBy(ctx context.Context, field, value string) ([]*Order, appError.IAppError)
	GetUserOrders(ctx context.Context, userId int) ([]*InfoForUser, appError.IAppError)
	Create(ctx context.Context, dto ShallowWeedProduct) (*int, appError.IAppError)
	ChangeStatus(ctx context.Context, statusId int) (*int, appError.IAppError)
}

type Service struct {
	validator *validator.Validate
	storage   IStorage
	logger    *logging.Logger
	cfg       *config.Config
}

func NewService(storage IStorage, validator *validator.Validate, logger *logging.Logger, cfg *config.Config) IService {
	return &Service{storage: storage, validator: validator, logger: logger, cfg: cfg}
}

func (s *Service) GetAll(ctx context.Context) ([]*Order, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetBy(ctx context.Context, field, value string) ([]*Order, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetUserOrders(ctx context.Context, userId int) ([]*InfoForUser, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Create(ctx context.Context, dto ShallowWeedProduct) (*int, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ChangeStatus(ctx context.Context, statusId int) (*int, appError.IAppError) {
	//TODO implement me
	panic("implement me")
}
