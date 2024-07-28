package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	d_entity "github.com/x-syaifullah-x/go-crud/src/internal/domain/repository/entity"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/repository/entity/role"
)

func TestCreateUser(t *testing.T) {
	t.Run("Error Username Exists", func(t *testing.T) {
		db, mock := NewMock(t)
		repo := NewAuthRepository(db)

		role := role.User()
		mockData := d_entity.NewUserData("name", "username", "email", "password", role)

		mock.ExpectExec(
			"INSERT INTO users (name, username, email, password, role) VALUES (?, ?, ?, ?, ?)",
		).WithArgs(
			mockData.GetName(),
			mockData.GetUsername(),
			mockData.GetEmail(),
			mockData.GetPassword(),
			role.Value(),
		).WillReturnError(fmt.Errorf("Error 1062 (23000): Duplicate entry '%s' for key 'username'", mockData.GetUsername()))
		res, err := repo.CreateUser(mockData)
		assert.NotNil(t, err)
		assert.Equal(t, d_entity.ID(0), res.GetID())
	})

	t.Run("Error Email Exists", func(t *testing.T) {
		db, mock := NewMock(t)
		repo := NewAuthRepository(db)

		role := role.User()
		mockData := d_entity.NewUserData("name", "username", "email", "password", role)

		mock.ExpectExec(
			"INSERT INTO users (name, username, email, password, role) VALUES (?, ?, ?, ?, ?)",
		).WithArgs(
			mockData.GetName(),
			mockData.GetUsername(),
			mockData.GetEmail(),
			mockData.GetPassword(),
			role.Value(),
		).WillReturnError(fmt.Errorf("Error 1062 (23000): Duplicate entry '%s' for key 'email'", mockData.GetEmail()))
		res, err := repo.CreateUser(mockData)
		assert.NotNil(t, err)
		assert.Equal(t, d_entity.ID(0), res.GetID())
	})

	t.Run("Error Unknown", func(t *testing.T) {
		db, mock := NewMock(t)
		repo := NewAuthRepository(db)

		role := role.User()
		mockData := d_entity.NewUserData("name", "username", "email", "password", role)

		mock.ExpectExec(
			"INSERT INTO users (name, username, email, password, role) VALUES (?, ?, ?, ?, ?)",
		).WithArgs(
			mockData.GetName(),
			mockData.GetUsername(),
			mockData.GetEmail(),
			mockData.GetPassword(),
			role.Value(),
		).WillReturnError(fmt.Errorf("Unknown"))
		res, err := repo.CreateUser(mockData)
		assert.NotNil(t, err)
		assert.Equal(t, d_entity.ID(0), res.GetID())
	})

	t.Run("Error LastInsertId", func(t *testing.T) {
		db, mock := NewMock(t)
		repo := NewAuthRepository(db)

		role := role.User()
		mockData := d_entity.NewUserData("name", "username", "email", "password", role)

		mock.ExpectExec(
			"INSERT INTO users (name, username, email, password, role) VALUES (?, ?, ?, ?, ?)",
		).WithArgs(
			mockData.GetName(),
			mockData.GetUsername(),
			mockData.GetEmail(),
			mockData.GetPassword(),
			role.Value(),
		).WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("LastInsertId")))
		res, err := repo.CreateUser(mockData)
		assert.NotNil(t, err)
		assert.Equal(t, d_entity.ID(0), res.GetID())
	})

	t.Run("Success", func(t *testing.T) {
		db, mock := NewMock(t)
		repo := NewAuthRepository(db)

		role := role.User()
		mockData := d_entity.NewUserData("name", "username", "email", "password", role)

		id := d_entity.ID(1)

		mock.ExpectExec(
			"INSERT INTO users (name, username, email, password, role) VALUES (?, ?, ?, ?, ?)",
		).WithArgs(
			mockData.GetName(),
			mockData.GetUsername(),
			mockData.GetEmail(),
			mockData.GetPassword(),
			role.Value(),
		).WillReturnResult(sqlmock.NewResult(int64(id), 1))
		res, err := repo.CreateUser(mockData)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, id, res.GetID())
	})
}

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
