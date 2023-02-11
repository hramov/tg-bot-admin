package order

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/config"
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type Storage interface {
	GetBy(ctx context.Context, field string, param any) (*Order, error)
	Get(ctx context.Context, limit int, lastId int) ([]*Order, error)
	Create(ctx context.Context, dto Order) (*int, error)
	ChangeStatus(ctx context.Context, orderId, statusId int) (*int, error)
	GetUserOrders(cts context.Context, userId int) ([]*InfoForUser, error)
}

type Service interface {
	GetAll(ctx context.Context, limit, offset int) ([]*Order, appError.IAppError)
	GetBy(ctx context.Context, field, value string) (*Order, appError.IAppError)
	GetUserOrders(ctx context.Context, userId int) ([]*InfoForUser, appError.IAppError)
	Create(ctx context.Context, dto Order) (*int, appError.IAppError)
	ChangeStatus(ctx context.Context, orderId, statusId int) (*int, appError.IAppError)
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

func (s *service) GetAll(ctx context.Context, limit, lastId int) ([]*Order, appError.IAppError) {
	products, err := s.storage.Get(ctx, limit, lastId)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return products, nil
}

func (s *service) GetBy(ctx context.Context, field, value string) (*Order, appError.IAppError) {
	order, err := s.storage.GetBy(ctx, field, value)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return order, nil
}

func (s *service) GetUserOrders(ctx context.Context, userId int) ([]*InfoForUser, appError.IAppError) {
	orders, err := s.storage.GetUserOrders(ctx, userId)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return orders, nil
}

func (s *service) Create(ctx context.Context, dto Order) (*int, appError.IAppError) {
	err := s.validator.Struct(dto)
	if err != nil {
		return nil, appError.ValidationError(err)
	}

	id, err := s.storage.Create(ctx, dto)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return id, nil
}

func (s *service) ChangeStatus(ctx context.Context, orderId, statusId int) (*int, appError.IAppError) {
	order, err := s.storage.ChangeStatus(ctx, orderId, statusId)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return order, nil
}
