package handler

import (
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/utils"
	"fr33d0mz/moneyflowx/utils/CustomError"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CheckWallet(c *gin.Context) {
	walletId := c.Query("id")

	if walletId == "" {
		response := utils.ErrorResponse("define id using query", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	wallet, err := h.service.Wallet.GetWalletByNumber(walletId)
	if err != nil {
		response := utils.ErrorResponse("wallet not found", http.StatusBadRequest, &CustomError.WalletNotFoundError{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, dto.FormatWallet(wallet))
}

func (h *Handler) GetBalance(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	input := &dto.WalletRequestBody{}
	input.UserID = user.ID
	wallet, err := h.service.Wallet.GetWalletByUserId(input)
	if err != nil {
		response := utils.ErrorResponse("show profile failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": wallet.UserID,
		"number":  wallet.Number,
		"balance": wallet.Balance,
		"type":    wallet.UserType,
	})
}
