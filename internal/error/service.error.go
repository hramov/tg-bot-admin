package appError

import "net/http"

func UserNotPresentInServiceError() IAppError {
	return &AppError{
		statusCode: http.StatusForbidden,
		message:    "User has no rights for the service",
	}
}

func CannotGenerateClientData() IAppError {
	return &AppError{
		statusCode: http.StatusInternalServerError,
		message:    "Cannot generate client id and client secret",
	}
}

func WrongClientDataError() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "Wrong client authorization data",
	}
}

func WrongUrlError() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "Wrong remote url",
	}
}

func NoSuchServiceError() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "No such service",
	}
}
