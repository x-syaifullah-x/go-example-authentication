package controller

import (
	"encoding/json"
	"net/http"

	"github.com/x-syaifullah-x/go-crud/src/internal/domain"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/controller"
	"github.com/x-syaifullah-x/go-crud/src/internal/domain/service"
	d_payload "github.com/x-syaifullah-x/go-crud/src/internal/domain/service/payload"
	"github.com/x-syaifullah-x/go-crud/src/internal/handler/controller/dto"
	"github.com/x-syaifullah-x/go-crud/src/internal/handler/controller/payload"
)

func NewAuthController(service service.AuthService) controller.AuthController {
	return &authController{service: service}
}

type authController struct {
	service service.AuthService
}

func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	payload, err := payload.MakeRegisterPayload(r.Body)

	if err != nil {
		sendResponse(w, dto.Response{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
			Error:         map[string]string{"message": err.Error()},
		})
		return
	}

	inputNotInvalid := payload.ValidateInputs()
	if len(inputNotInvalid) > 0 {
		sendResponse(w, dto.Response{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
			Error:         inputNotInvalid,
		})
		return
	}

	result, err := c.service.Register(
		d_payload.NewRegisterPayload(
			payload.Name,
			payload.Username,
			payload.Email,
			payload.Password,
		),
	)
	if err != nil {
		statusCode := 0
		switch err.(type) {
		case *domain.ErrEmailAlreadyExists:
			statusCode = http.StatusConflict
		case *domain.ErrUsernameAlreadyExists:
			statusCode = http.StatusConflict
		case *domain.ErrDatabase:
			statusCode = http.StatusInternalServerError
		default:
			statusCode = http.StatusBadRequest
		}
		sendResponse(w, dto.Response{
			StatusCode:    statusCode,
			StatusMessage: http.StatusText(statusCode),
			Error: map[string]string{
				"message": err.Error(),
			},
		})
		return
	}

	sendResponse(w, dto.Response{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data: map[string]string{
			"id": result.GetName(),
		},
	})
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func sendResponse(w http.ResponseWriter, r dto.Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	json.NewEncoder(w).Encode(r)
}
