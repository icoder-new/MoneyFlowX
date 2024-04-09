package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/icoder-new/MoneyFlowX/models"
	"github.com/icoder-new/MoneyFlowX/pkg/dto"
	"github.com/icoder-new/MoneyFlowX/utils"
	"github.com/icoder-new/MoneyFlowX/utils/CustomError"
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
