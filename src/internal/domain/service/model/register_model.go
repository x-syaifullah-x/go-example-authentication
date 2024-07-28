package model

func NewRegisterModel(name string) RegisterModel {
	return RegisterModel{name: name}
}

type RegisterModel struct {
	name string
}

func (r RegisterModel) GetName() string {
	return r.name
}
