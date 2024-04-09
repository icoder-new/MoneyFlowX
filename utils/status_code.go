package utils

import (
	"github.com/icoder-new/MoneyFlowX/utils/CustomError"
	"net/http"
)

func GetStatusCode(err error) int {
	var statusCode int = http.StatusInternalServerError

	if _, ok := err.(*CustomError.NotValidEmailError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*CustomError.UserAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*CustomError.IncorrectCredentialsError); ok {
		statusCode = http.StatusUnauthorized
	} else if _, ok := err.(*CustomError.UserNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*CustomError.PasswordNotSameError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*CustomError.ResetTokenNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*CustomError.InsufficientBalanceError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*CustomError.WalletNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*CustomError.WalletAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*CustomError.TransferToSameWalletError); ok {
		statusCode = http.StatusBadRequest
	}

	return statusCode
}
