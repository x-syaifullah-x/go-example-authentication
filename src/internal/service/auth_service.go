package service

import (
	"fmt"

	"github.com/x-syaifullah-x/go-crud/src/internal/domain/repository"
	d_entity "github.com/x-syaifullah-x/go-crud/src/internal/domain/repository/entity"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/repository/entity/role"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/service"
	d_model "github.com/x-syaifullah-x/go-crud/src/internal/domain/service/model"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/service/payload"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(repository repository.AuthRepository) service.AuthService {
	return &authService{repository: repository}
}

type authService struct {
	repository repository.AuthRepository
}

func (s *authService) Register(payload payload.RegisterPayload) (d_model.RegisterModel, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(payload.GetPassword()),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return d_model.RegisterModel{}, err
	}

	result, err := s.repository.CreateUser(d_entity.NewUserData(
		payload.GetName(),
		payload.GetUsername(),
		payload.GetEmail(),
		string(passwordHash),
		role.User(),
	))

	if err != nil {
		return d_model.RegisterModel{}, err
	}
	name := fmt.Sprint(result.GetID())
	return d_model.NewRegisterModel(name), nil
}

func (s *authService) Login(payload payload.LoginPayload) (d_model.LoginModel, error) {
	panic("NOT IMPLEMENTED")
}
