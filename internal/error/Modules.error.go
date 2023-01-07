package appError

import "net/http"

func DatabaseError(err error) IAppError {
	return &AppError{
		statusCode: http.StatusInternalServerError,
		message:    err.Error(),
	}
}

func CannotSendMail(err error) IAppError {
	return &AppError{
		statusCode: http.StatusInternalServerError,
		message:    "Cannot send mail: " + err.Error(),
	}
}

func NoUserFoundError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "No user found",
	}
}

func NoClientIdFoundError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "No client id found",
	}
}

func NoUserIpFoundError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "No user ip found",
	}
}

func NoLoginFoundError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "No login found",
	}
}

func NoPasswordFoundError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "No password found",
	}
}

func NoAuthCodeError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "No auth code found",
	}
}

func NoIdFoundError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "No id found",
	}
}

func WrongIdFormatError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "Wrong id format",
	}
}

func CannotGetIdError() IAppError {
	return &AppError{
		statusCode: http.StatusBadRequest,
		message:    "Cannot get id",
	}
}
