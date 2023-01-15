package product

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/config"
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type Storage interface {
	GetBy(ctx context.Context, field string, param any) (*Product, error)
	Get(ctx context.Context, limit int, lastId int) ([]*Product, error)
	Create(ctx context.Context, dto InputWeedProduct) (*int, error)
	Update(ctx context.Context, dto InputWeedProduct) (*int, error)
	Delete(ctx context.Context, id int) (*int, error)
}

type Service interface {
	GetAll(ctx context.Context, limit int, lastId int) ([]*Product, appError.IAppError)
	GetBy(ctx context.Context, field, value string) (*Product, appError.IAppError)
	Create(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError)
	Update(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError)
	Delete(ctx context.Context, id int) (*int, appError.IAppError)
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

func (s *service) GetAll(ctx context.Context, limit, lastId int) ([]*Product, appError.IAppError) {
	products, err := s.storage.Get(ctx, limit, lastId)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return products, nil
}

func (s *service) GetBy(ctx context.Context, field, value string) (*Product, appError.IAppError) {
	product, err := s.storage.GetBy(ctx, field, value)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return product, nil
}

func (s *service) Create(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError) {
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

func (s *service) Update(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError) {
	err := s.validator.Struct(dto)
	if err != nil {
		return nil, appError.ValidationError(err)
	}

	id, err := s.storage.Update(ctx, dto)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return id, nil
}

func (s *service) Delete(ctx context.Context, id int) (*int, appError.IAppError) {
	deletedId, err := s.storage.Delete(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return deletedId, nil
}
