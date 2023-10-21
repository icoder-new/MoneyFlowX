package dto

import (
	"fr33d0mz/moneyflowx/models"
	"strings"
	"time"
)

type TopUpRequestBody struct {
	Amount         float64 `json:"amount" binding:"required,min=1,max=5000"`
	SourceOfFundID string  `json:"source_of_fund_id" binding:"required"`
	User           *models.User
}

type TransferRequestBody struct {
	Amount       float64 `json:"amount" binding:"required,min=1,max=5000"`
	WalletNumber string  `json:"wallet_number" binding:"required"`
	Comment      string  `json:"comment"`
	User         *models.User
}

type TopUpResponse struct {
	ID            string    `json:"uuid"`
	SourceOfFund  string    `json:"source_of_fund"`
	Amount        float64   `json:"amount"`
	WalletBalance float64   `json:"balance"`
	Comment       string    `json:"comment"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at" time_format:"2006-01-02"`
	UpdatedAt     time.Time `json:"updated_at" time_format:"2006-01-02"`
}

type TransactionRequestQuery struct {
	Search string `form:"s"`
	SortBy string `form:"sortBy"`
	Sort   string `form:"sort"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type Destination struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

type TransferResponse struct {
	ID            string      `json:"uuid"`
	Destination   Destination `json:"destination"`
	Amount        float64     `json:"amount"`
	WalletBalance float64     `json:"balance"`
	Comment       string      `json:"comment"`
	Type          string      `json:"type"`
	CreatedAt     time.Time   `json:"created_at" time_format:"2006-01-02"`
	UpdatedAt     time.Time   `json:"updated_at" time_format:"2006-01-02"`
}

type TransactionResponse struct {
	ID           string      `json:"uuid"`
	SourceOfFund string      `json:"source_of_fund"`
	Destination  Destination `json:"destination"`
	Amount       float64     `json:"amount"`
	Comment      string      `json:"comment"`
	Type         string      `json:"Type"`
	CreatedAt    time.Time   `json:"created_at" time_format:"2006-01-02"`
	UpdatedAt    time.Time   `json:"updated_at" time_format:"2006-01-02"`
}

func FormatTopUp(transaction *models.Transaction) TopUpResponse {
	return TopUpResponse{
		ID:            transaction.ID,
		SourceOfFund:  transaction.SourceOfFund.Name,
		Amount:        transaction.Amount,
		WalletBalance: transaction.Wallet.Balance,
		Comment:       transaction.Comment,
		Type:          transaction.Type,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func FormatTransfer(transaction *models.Transaction) TransferResponse {
	return TransferResponse{
		ID:            transaction.ID,
		Destination:   Destination{Name: transaction.Wallet.User.Username, Number: transaction.Wallet.Number},
		Amount:        transaction.Amount,
		WalletBalance: transaction.Wallet.Balance,
		Comment:       transaction.Comment,
		Type:          transaction.Type,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func FormatTransaction(transaction *models.Transaction) TransactionResponse {
	var sourceOfFund string
	if transaction.SourceOfFund != nil {
		sourceOfFund = transaction.SourceOfFund.Name
	}
	return TransactionResponse{
		ID:           transaction.ID,
		SourceOfFund: sourceOfFund,
		Destination:  Destination{Name: transaction.Wallet.User.Username, Number: transaction.Wallet.Number},
		Amount:       transaction.Amount,
		Comment:      transaction.Comment,
		Type:         transaction.Type,
		CreatedAt:    transaction.CreatedAt,
		UpdatedAt:    transaction.UpdatedAt,
	}
}

func FormatTransactions(transactions []*models.Transaction) []TransactionResponse {
	formattedTransactions := []TransactionResponse{}
	for _, transaction := range transactions {
		formattedBook := FormatTransaction(transaction)
		formattedTransactions = append(formattedTransactions, formattedBook)
	}
	return formattedTransactions
}

func FormatQuery(query *TransactionRequestQuery) *TransactionRequestQuery {
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	query.SortBy = strings.ToLower(query.SortBy)
	if query.SortBy == "date" {
		query.SortBy = "updated_at"
	} else if query.SortBy == "to" {
		query.SortBy = "destination_id"
	} else if query.SortBy != "amount" {
		query.SortBy = "updated_at"
	}

	query.Sort = strings.ToUpper(query.Sort)
	if query.Sort != "ASC" {
		query.Sort = "DESC"
	}

	return query
}
