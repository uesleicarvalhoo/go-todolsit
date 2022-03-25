package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/internal/config"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/pkg/auth"
	"github.com/uesleicarvalhoo/go-todolist/pkg/utils"
)

type Service struct {
	repository Repository
}

func NewService(userRepo Repository) *Service {
	return &Service{
		repository: userRepo,
	}
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) SiginUp(ctx context.Context, payload entity.SiginUp) (*entity.User, error) {
	if u, _ := s.repository.GetByEmail(ctx, payload.Email); u.Email == payload.Email {
		return nil, errors.New(utils.ErrEmailDuplicated)
	}

	user, err := entity.NewUser(payload.Name, payload.Email, payload.Password)
	if err != nil {
		return nil, err
	}

	s.repository.Create(ctx, user)
	return user, nil
}

func (s *Service) Login(ctx context.Context, payload entity.Login) (entity.LoginResponse, error) {
	user, err := s.repository.GetByEmail(ctx, payload.Email)
	if err != nil {
		return entity.LoginResponse{
			Message: utils.ErrEmailNotFound,
		}, err
	}

	if !user.ValidatePassword(payload.Password) {
		return entity.LoginResponse{
			Message: utils.ErrPasswordInvalid,
		}, errors.New("Unauthorized")
	}

	token, err := auth.GenerateJWT(config.GetEnv().SecretKey, user.Id)
	if err != nil {
		return entity.LoginResponse{
			Message: utils.ErrGenerateToken,
		}, errors.New("Error on generate token")
	}
	return entity.LoginResponse{
		Token:   token,
		Message: "Success",
	}, nil
}
