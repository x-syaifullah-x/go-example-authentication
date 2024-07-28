package handler

import (
	"fmt"
	"net/http"

	"github.com/x-syaifullah-x/go-crud/src/external/database"
	"github.com/x-syaifullah-x/go-crud/src/internal/config"
	"github.com/x-syaifullah-x/go-crud/src/internal/handler/controller"
	"github.com/x-syaifullah-x/go-crud/src/internal/middleware"
	"github.com/x-syaifullah-x/go-crud/src/internal/repository"
	"github.com/x-syaifullah-x/go-crud/src/internal/service"
	"github.com/x-syaifullah-x/go-crud/src/pkg/logger"
)

func NewHandler(server *http.ServeMux) http.Handler {
	authMiddleware := middleware.NewAuthMiddleware(server)

	db := database.Instance()
	repo := repository.NewAuthRepository(db)
	service := service.NewAuthService(repo)
	controller := controller.NewAuthController(service)

	server.HandleFunc("/api/auth/register", controller.Register)
	server.HandleFunc("/api/auth/login", controller.Login)

	logger.Print(fmt.Sprint("Listen on port ", config.GetConfig().App.Port))
	
	return authMiddleware
}
