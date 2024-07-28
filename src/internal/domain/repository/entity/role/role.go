package role

type role string

type Role struct {
	value role
}

var admin = Role{"admin"}
var user = Role{"user"}
var guest = Role{"guest"}

func Admin() Role {
	return admin
}

func User() Role {
	return user
}

func Guest() Role {
	return guest
}

func (r Role) Value() string {
	return string(r.value)
}
