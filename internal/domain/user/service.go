package user

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/config"
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/hramov/tg-bot-admin/pkg/jwt"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type IStorage interface {
	GetBy(ctx context.Context, field string, param any) (*User, error)
	GetById(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Get(ctx context.Context, limit, offset int) ([]*User, error)
	Create(ctx context.Context, dto *CreateDto) (*int, error)
	Update(ctx context.Context, dto *UpdateDto) (*int, error)
	Delete(ctx context.Context, id int) (*int, error)
}

type IService interface {
	GetAll(ctx context.Context, limit, offset int) ([]*User, appError.IAppError)
	GetById(ctx context.Context, id int) (*User, appError.IAppError)
	Create(ctx context.Context, dto *CreateDto) (*int, appError.IAppError)
	Update(ctx context.Context, dto *UpdateDto) (*int, appError.IAppError)
	Delete(ctx context.Context, id int) (*int, appError.IAppError)
	Login(ctx context.Context, dto *LoginDto) (*LoginResponseDto, appError.IAppError)
	Refresh(ctx context.Context, dto *LoginResponseDto) (*LoginResponseDto, appError.IAppError)
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

func (s *Service) Login(ctx context.Context, dto *LoginDto) (*LoginResponseDto, appError.IAppError) {
	err := s.validator.Struct(dto)
	if err != nil {
		return nil, appError.ValidationError(err)
	}
	user, err := s.storage.GetBy(ctx, "email", dto.Email)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	valid := jwt.CheckPassword(dto.Password, user.Password)
	if !valid {
		return nil, appError.LoginOrPasswordIncorrectError()
	}
	at, rt, err := jwt.CreateToken(user.Id, s.cfg.Jwt.AccessSecret, s.cfg.Jwt.RefreshSecret, s.cfg.Jwt.AccessTtl, s.cfg.Jwt.RefreshTtl)
	if err != nil {
		return nil, appError.CreateTokenError()
	}
	return &LoginResponseDto{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, dto *LoginResponseDto) (*LoginResponseDto, appError.IAppError) {
	id, err := jwt.CheckRefreshToken(dto.RefreshToken, s.cfg.Jwt.RefreshSecret)
	if err != nil {
		return nil, appError.RefreshTokenIsInvalidError()
	}
	user, err := s.storage.GetBy(ctx, "id", id)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	if user.Id == 0 {
		return nil, appError.NoUserFoundError()
	}
	at, rt, err := jwt.CreateToken(user.Id, s.cfg.Jwt.AccessSecret, s.cfg.Jwt.RefreshSecret, s.cfg.Jwt.AccessTtl, s.cfg.Jwt.RefreshTtl)
	if err != nil {
		return nil, appError.CreateTokenError()
	}
	return &LoginResponseDto{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func (s *Service) GetAll(ctx context.Context, limit, offset int) ([]*User, appError.IAppError) {
	users, err := s.storage.Get(ctx, limit, offset)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	return users, nil
}

func (s *Service) GetById(ctx context.Context, id int) (*User, appError.IAppError) {
	user, err := s.storage.GetBy(ctx, "id", id)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	return user, nil
}

func (s *Service) Create(ctx context.Context, dto *CreateDto) (*int, appError.IAppError) {
	err := s.validator.Struct(dto)
	if err != nil {
		return nil, appError.ValidationError(err)
	}
	dto.Password, err = jwt.CreateHashedPassword(dto.Password)
	if err != nil {
		return nil, appError.InternalServerError()
	}
	id, err := s.storage.Create(ctx, dto)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	return id, nil
}

func (s *Service) Update(ctx context.Context, dto *UpdateDto) (*int, appError.IAppError) {
	err := s.validator.Struct(dto)
	if err != nil {
		return nil, appError.ValidationError(err)
	}

	if dto.Password != "" {
		dto.Password, err = jwt.CreateHashedPassword(dto.Password)
	}

	if err != nil {
		return nil, appError.InternalServerError()
	}
	id, err := s.storage.Update(ctx, dto)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	return id, nil
}

func (s *Service) Delete(ctx context.Context, id int) (*int, appError.IAppError) {
	deletedId, err := s.storage.Delete(ctx, id)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	return deletedId, nil
}
