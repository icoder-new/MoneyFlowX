package handler

import (
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/utils"
	"github.com/gin-gonic/gin"
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
