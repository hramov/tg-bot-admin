package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type IStorage interface {
}

type IService interface {
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
