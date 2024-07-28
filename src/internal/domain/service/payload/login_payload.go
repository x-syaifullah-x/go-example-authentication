package payload

func NewLoginPayload(
	username string,
	email string,
	password string,
) LoginPayload {
	return LoginPayload{
		username: username,
		email:    email,
		password: password,
	}
}

type LoginPayload struct {
	username string
	email    string
	password string
}

func (l LoginPayload) GetUsername() string {
	return l.username
}
func (l LoginPayload) GetEmail() string {
	return l.email
}
func (l LoginPayload) GetPassword() string {
	return l.password
}
