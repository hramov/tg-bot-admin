package utils

import appError "github.com/hramov/tg-bot-admin/internal/error"

type AppResponse[T comparable] struct {
	Status bool
	Error  appError.IAppError
	Data   *T
}

func CreateAppResponse[T comparable](status bool, err appError.IAppError, data *T) *AppResponse[T] {
	return &AppResponse[T]{
		Status: status,
		Error:  err,
		Data:   data,
	}
}

func EqualSlice(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
