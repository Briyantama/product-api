package constants

import (
	"fmt"
	"net/http"
)

const (
	ERRBADREQUEST = iota
	ERRUNAUTHORIZED
	ERRNOTFOUND
	ERRCONFLICT
	DEFAULT
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}

func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func ToHTTPStatus(code int) int {
	switch code {
	case ERRCONFLICT:
		return http.StatusConflict
	case ERRNOTFOUND:
		return http.StatusNotFound
	case ERRUNAUTHORIZED:
		return http.StatusUnauthorized
	case ERRBADREQUEST:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

const (
	// User errors
	ErrUserAlreadyExists    = "user already exists"
	ErrUserRegistrationFail = "user registration failed"
	ErrUserAuthentication   = "user authentication failed"
	ErrUserNotFound         = "user not found"
	ErrUserInvalid          = "email or password invalid"

	// Vendor errors
	ErrVendorAlreadyExists    = "vendor already exists"
	ErrVendorRegistrationFail = "vendor registration failed"
	ErrVendorNotFound         = "vendor not found"
	ErrVendorAccessDenied     = "vendor access denied"
	ErrGetVendorByUserIDFail  = "failed to get vendor by user ID"

	// Product errors
	ErrProductCreationFail   = "product creation failed"
	ErrProductUpdateFail     = "product update failed"
	ErrProductDeleteFail     = "product deletion failed"
	ErrProductNotFound       = "product not found"
	ErrProductAccessDenied   = "you are not allowed to access this product"
	ErrProductInvalidPayload = "invalid product payload"

	// General/common
	ErrInvalidID          = "invalid ID"
	ErrInvalidRequestData = "invalid request data"
	ErrUnauthorized       = "unauthorized"
	ErrDatabaseOperation  = "database operation failed"
	ErrTransactionFail    = "transaction failed"
	Err                   = "error"
)
