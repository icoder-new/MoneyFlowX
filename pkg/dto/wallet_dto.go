package dto

import "fr33d0mz/moneyflowx/models"

type WalletRequestBody struct {
	UserID string `json:"name" binding:"required"`
}

type WalletResponse struct {
	ID      string  `json:"uuid"`
	Number  string  `json:"number"`
	Balance float64 `json:"balance"`
}

func FormatWallet(wallet *models.Wallet) WalletResponse {
	return WalletResponse{
		ID:      wallet.ID,
		Number:  wallet.Number,
		Balance: wallet.Balance,
	}
}
