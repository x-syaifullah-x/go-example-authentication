package domain

import "fmt"

type ErrEmailAlreadyExists struct {
	Email string
}

func (e ErrEmailAlreadyExists) Error() string {
	return fmt.Sprintf("The email %s already exists.", e.Email)
}

type ErrUsernameAlreadyExists struct {
	Username string
}

func (e ErrUsernameAlreadyExists) Error() string {
	return fmt.Sprintf("The username %s already exists.", e.Username)
}

type ErrDatabase struct {
	Message string
}

func (e ErrDatabase) Error() string {
	return e.Message
}
