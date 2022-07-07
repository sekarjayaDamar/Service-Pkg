package utilities

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (appError AppError) AsMessage() *AppError {
	return &AppError{
		Message: appError.Message,
	}
}

func NewNotFoundError(Message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: Message,
	}
}

func NewUnexpectedError(Message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: Message,
	}
}

func NewBadRequest(Message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: Message,
	}
}

func NewValidationError(Message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: Message,
	}
}

func NewLoadConfigError(Message string) *AppError {
	return &AppError{
		Code:    http.StatusNoContent,
		Message: Message,
	}
}

func NewUnauthorizedError(Message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: Message,
	}
}
