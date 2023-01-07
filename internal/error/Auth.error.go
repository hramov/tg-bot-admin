package appError

import (
	"fmt"
	"net/http"
)

func CSRFForbiddenError() IAppError {
	return &AppError{
		statusCode: http.StatusForbidden,
		message:    "CSRF token forbidden",
	}
}

func AuthCodeError() IAppError {
	return &AppError{
		statusCode: http.StatusForbidden,
		message:    "Authorization code is invalid",
	}
}

func WrongClientSecretError() IAppError {
	return &AppError{
		statusCode: http.StatusForbidden,
		message:    "Client id or client secret is wrong",
	}
}

func NoCredentialsError() IAppError {
	return &AppError{
		statusCode: http.StatusForbidden,
		message:    "Username or password not passed",
	}
}

func WrongCredFormat(field string) IAppError {
	return &AppError{
		statusCode: http.StatusForbidden,
		message:    fmt.Sprintf("Field %s has wrong format", field),
	}
}
