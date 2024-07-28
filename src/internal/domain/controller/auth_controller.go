package controller

import (
	"net/http"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
