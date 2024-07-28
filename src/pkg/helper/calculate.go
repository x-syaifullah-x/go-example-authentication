package helper

import (
	"unsafe"

	"github.com/x-syaifullah-x/go-crud/src/internal/domain/service/model"
)

func CalculateSizeToID(data map[string]uint32) uintptr {
	var size uintptr
	for key, value := range data {
		size += unsafe.Sizeof(key)
		size += uintptr(len(key))
		size += unsafe.Sizeof(value)
	}
	return size
}

func CalculateSizeCach(data map[uint32]model.UserModel) uintptr {
	var size uintptr
	for key, value := range data {
		size += unsafe.Sizeof(key) + unsafe.Sizeof(value)
		size += uintptr(len(value.GetName()))
		size += uintptr(len(value.GetEmail()))
	}
	return size
}
