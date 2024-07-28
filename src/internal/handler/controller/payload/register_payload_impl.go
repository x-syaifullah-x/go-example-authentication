package payload

import (
	"encoding/json"
	"io"
)

func MakeRegisterPayloada() (*RegisterPayload, error) {
	payload := &RegisterPayload{}
	return payload, nil
}

func MakeRegisterPayload(body io.ReadCloser) (*RegisterPayload, error) {
	payload := &RegisterPayload{}
	err := json.NewDecoder(body).Decode(payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

type RegisterPayload struct {
	Name            string `json:"name,omitempty"`
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

func (p *RegisterPayload) ValidateInputs() map[string]string {
	fieldErrors := make(map[string]string)
	if len(p.Name) == 0 {
		fieldErrors["name"] = "name is reqired"
	}

	if len(p.Username) == 0 {
		fieldErrors["username"] = "username is reqired"
	}

	if len(p.Email) == 0 {
		fieldErrors["email"] = "email is reqired"
	}

	if len(p.Password) == 0 {
		fieldErrors["password"] = "password is reqired"
	} else {
		if len(p.ConfirmPassword) == 0 {
			fieldErrors["confirm_password"] = "confirm password is reqired"
		} else {
			if p.Password != p.ConfirmPassword {
				fieldErrors["message"] = "Password and confirm password do not match"
			}
		}
	}
	return fieldErrors
}
