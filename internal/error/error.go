package appError

import "net/http"

type IAppError interface {
	error
	Status() int
}

type AppError struct {
	statusCode int
	message    string
	code       string
}

func (a AppError) Error() string {
	return a.message
}

func (a AppError) Status() int {
	return a.statusCode
}

func InternalServerError() IAppError {
	return &AppError{
		statusCode: http.StatusInternalServerError,
		message:    "Cannot serve the request",
	}
}

func NotImplementedError() IAppError {
	return &AppError{
		statusCode: http.StatusInternalServerError,
		message:    "Method not implement",
	}
}

func ValidationError(err error) IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    err.Error(),
	}
}
