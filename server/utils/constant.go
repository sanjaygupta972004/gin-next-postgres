package utils

import "errors"

var (
	// General Errors
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized access")
	ErrForbidden      = errors.New("forbidden access")
	ErrNotFound       = errors.New("resource not found")
	ErrConflict       = errors.New("resource conflict")
	ErrValidation     = errors.New("validation error")

	// Authentication Errors
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("token has expired")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrAccountLocked      = errors.New("account is locked")

	// Database Errors
	ErrDBConnection   = errors.New("database connection error")
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
)
