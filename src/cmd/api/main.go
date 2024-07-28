package main

import (
	"fmt"
	"net/http"

	// _ "net/http/pprof"

	"github.com/x-syaifullah-x/go-crud/src/external/database"
	"github.com/x-syaifullah-x/go-crud/src/internal/config"
	"github.com/x-syaifullah-x/go-crud/src/internal/handler"
	"github.com/x-syaifullah-x/go-crud/src/pkg/logger"
)

func main() {
	// go func() {
	// 	logger.Fatal(http.ListenAndServe("localhost:6060", nil))
	// }()

	err := config.LoadConfig("src/cmd/api/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	defer database.Instance().Close()

	err = http.ListenAndServe(
		fmt.Sprintf(":%s", config.GetConfig().App.Port),
		handler.NewHandler(http.NewServeMux()),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
