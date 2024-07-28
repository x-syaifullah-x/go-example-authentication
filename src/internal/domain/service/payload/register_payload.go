package payload

func NewRegisterPayload(
	name string,
	username string,
	email string,
	password string,
) RegisterPayload {
	return RegisterPayload{
		name:     name,
		username: username,
		email:    email,
		password: password,
	}
}

type RegisterPayload struct {
	name     string
	username string
	email    string
	password string
}

func (r RegisterPayload) GetName() string {
	return r.name
}

func (r RegisterPayload) GetUsername() string {
	return r.username
}

func (r RegisterPayload) GetEmail() string {
	return r.email
}

func (r RegisterPayload) GetPassword() string {
	return r.password
}
