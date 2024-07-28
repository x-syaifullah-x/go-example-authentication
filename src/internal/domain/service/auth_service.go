package service

import (
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/service/model"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/service/payload"
)

type AuthService interface {
	Register(payload payload.RegisterPayload) (model.RegisterModel, error)
	Login(payload payload.LoginPayload) (model.LoginModel, error)
}
