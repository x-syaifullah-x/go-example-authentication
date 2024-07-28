package repository

import "github.com/x-syaifullah-x/go-crud/src/internal/domain/repository/entity"

type AuthRepository interface {
	CreateUser(data entity.UserData) (entity.CreateResult, error)
}
