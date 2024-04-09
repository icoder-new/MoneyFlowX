package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/icoder-new/MoneyFlowX/logger"
	"github.com/icoder-new/MoneyFlowX/models"
	"github.com/icoder-new/MoneyFlowX/pkg/dto"
	"github.com/icoder-new/MoneyFlowX/utils"
	"net/http"
)

func (h *Handler) Profile(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	input := &dto.WalletRequestBody{}
	input.UserID = user.ID
	wallet, err := h.service.Wallet.GetWalletByUserId(input)
	if err != nil {
		response := utils.ErrorResponse("show profile failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedUser := dto.FormatUserDetail(user, wallet)
	response := utils.SuccessResponse("show profile success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) SendToken(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user.Type == "identified" {
		c.AbortWithStatusJSON(http.StatusOK,
			utils.SuccessResponse("user already identified", http.StatusOK, user))
		return
	}

	if err := h.service.JWT.SendVerificationToken(user); err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("sending token failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("verify link was send", http.StatusOK, nil))
}

func (h *Handler) VerifyUser(c *gin.Context) {
	_token := c.Param("verifyToken")
	logger.Debug.Println("token: ", _token)

	user, wallet, err := h.service.JWT.VerifyUser(_token)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("verifying user failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedResponse := dto.FormatUserDetail(user, wallet)
	c.JSON(http.StatusOK, utils.SuccessResponse("user identified", http.StatusOK, formattedResponse))
}
