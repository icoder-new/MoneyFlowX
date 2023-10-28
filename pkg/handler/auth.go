package handler

import (
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	input := &dto.RegisterRequestBody{}

	if err := c.ShouldBindJSON(input); err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("register failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.service.User.CreateUser(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("register failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	inputWallet := &dto.WalletRequestBody{}
	inputWallet.UserID = newUser.ID
	newWallet, err := h.service.Wallet.CreateWallet(inputWallet)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("register failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	token, err := h.service.JWT.GenerateToken(newUser.ID)
	if err != nil {
		response := utils.ErrorResponse("register failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(newUser, newWallet, token)
	response := utils.SuccessResponse("register success", http.StatusCreated, formattedLogin)
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) Login(c *gin.Context) {
	input := &dto.LoginRequestBody{}

	if err := c.ShouldBindJSON(input); err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("login failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.service.Authorization.Attempt(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("login failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	inputWallet := &dto.WalletRequestBody{}
	inputWallet.UserID = loggedInUser.ID
	wallet, err := h.service.Wallet.GetWalletByUserId(inputWallet)
	if err != nil {
		response := utils.ErrorResponse("login failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := h.service.JWT.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := utils.ErrorResponse("login failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(loggedInUser, wallet, token)
	response := utils.SuccessResponse("login success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	input := &dto.ForgotPasswordRequestBody{}

	if err := c.ShouldBindJSON(input); err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("forgot password failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	passwordReset, err := h.service.Authorization.ForgotPass(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("forgot password failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedPasswordReset := dto.FormatForgotPassword(passwordReset)
	response := utils.SuccessResponse("forgot password success", http.StatusOK, formattedPasswordReset)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) ResetPassword(c *gin.Context) {
	input := &dto.ResetPasswordRequestBody{}

	if err := c.ShouldBindJSON(input); err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("reset password failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	passwordReset, err := h.service.Authorization.ResetPass(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("reset password failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedUser := dto.FormatUser(&passwordReset.User)
	response := utils.SuccessResponse("reset password success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}
