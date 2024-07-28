package jwt

import (
	"testing"

	"github.com/x-syaifullah-x/go-crud/src/pkg/logger"
)

func TestGenerateSecretKey(t *testing.T) {
	logger.Print(generateSecretKey())
}
