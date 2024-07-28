package payload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRegisterPayload(t *testing.T) {
	a, _ := MakeRegisterPayloada()
	assert.Nil(t, a.ConfirmPassword)
}
