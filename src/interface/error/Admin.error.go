package appError

import "net/http"

func RefreshTokenIsInvalidError() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "Wrong refresh token format",
	}
}

func TokenDataParsingError() IAppError {
	return &AppError{
		statusCode: http.StatusInternalServerError,
		message:    "Cannot parse data from token",
	}
}

func CreateTokenError() IAppError {
	return &AppError{
		statusCode: http.StatusInternalServerError,
		message:    "Cannot create token pair",
	}
}

func AdminNotFound() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "Admin not found",
	}
}

func LoginOrPasswordIncorrectError() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "Login or password is incorrect",
	}
}

func WrongIpAddressError() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "Wrong IP-address",
	}
}

func WrongCurrentPasswordError() IAppError {
	return &AppError{
		statusCode: http.StatusUnauthorized,
		message:    "Wrong current password",
	}
}
