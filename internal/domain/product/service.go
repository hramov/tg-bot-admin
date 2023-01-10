package product

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/config"
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type IStorage interface {
	GetBy(ctx context.Context, field string, param any) (*Product, error)
	Get(ctx context.Context) ([]*Product, error)
	Create(ctx context.Context, dto InputWeedProduct) (*int, error)
	Update(ctx context.Context, dto InputWeedProduct) (*int, error)
	Delete(ctx context.Context, id int) (*int, error)
}

type IService interface {
	GetAll(ctx context.Context) ([]*Product, appError.IAppError)
	GetBy(ctx context.Context, field, value string) (*Product, appError.IAppError)
	Create(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError)
	Update(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError)
	Delete(ctx context.Context, id int) (*int, appError.IAppError)
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

func (s *Service) GetAll(ctx context.Context) ([]*Product, appError.IAppError) {
	products, err := s.storage.Get(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return products, nil
}

func (s *Service) GetBy(ctx context.Context, field, value string) (*Product, appError.IAppError) {
	product, err := s.storage.GetBy(ctx, field, value)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return product, nil
}

func (s *Service) Create(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError) {
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

func (s *Service) Update(ctx context.Context, dto InputWeedProduct) (*int, appError.IAppError) {
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

func (s *Service) Delete(ctx context.Context, id int) (*int, appError.IAppError) {
	deletedId, err := s.storage.Delete(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, appError.DatabaseError(err)
	}
	return deletedId, nil
}
