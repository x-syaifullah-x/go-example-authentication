package entity

import "github.com/x-syaifullah-x/go-crud/src/internal/domain/repository/entity/role"

func NewUserData(
	name string,
	username string,
	email string,
	password string,
	role role.Role,
) UserData {
	return UserData{
		name:     name,
		username: username,
		email:    email,
		password: password,
		role:     role,
	}
}

type UserData struct {
	name     string
	username string
	email    string
	password string
	role     role.Role
}

func (u UserData) GetName() string {
	return u.name
}
func (u UserData) GetUsername() string {
	return u.username
}
func (u UserData) GetEmail() string {
	return u.email
}
func (u UserData) GetPassword() string {
	return u.password
}
func (u UserData) GetRole() role.Role {
	return u.role
}
