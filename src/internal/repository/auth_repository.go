package repository

import (
	"database/sql"
	"fmt"

	"github.com/x-syaifullah-x/go-crud/src/internal/domain"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/repository"
	d_entity "github.com/x-syaifullah-x/go-crud/src/internal/domain/repository/entity"
)

func NewAuthRepository(db *sql.DB) repository.AuthRepository {
	return authRepository{db: db}
}

type authRepository struct {
	db *sql.DB
}

func (r authRepository) CreateUser(data d_entity.UserData) (d_entity.CreateResult, error) {
	execInsert, err := r.db.Exec(
		"INSERT INTO users (name, username, email, password, role) VALUES (?, ?, ?, ?, ?)",
		data.GetName(),
		data.GetUsername(),
		data.GetEmail(),
		data.GetPassword(),
		data.GetRole().Value(),
	)
	if err != nil {
		if err.Error() == fmt.Sprintf("Error 1062 (23000): Duplicate entry '%s' for key 'username'", data.GetUsername()) {
			return d_entity.CreateResult{}, domain.ErrUsernameAlreadyExists{Username: data.GetUsername()}
		}
		if err.Error() == fmt.Sprintf("Error 1062 (23000): Duplicate entry '%s' for key 'email'", data.GetEmail()) {
			return d_entity.CreateResult{}, domain.ErrEmailAlreadyExists{Email: data.GetEmail()}
		}
		return d_entity.CreateResult{}, domain.ErrDatabase{Message: err.Error()}
	}
	id, err := execInsert.LastInsertId()
	if err != nil {
		return d_entity.CreateResult{}, domain.ErrDatabase{Message: err.Error()}
	}

	return d_entity.NewCreateResult(d_entity.ID(id)), nil
}
