package service

import (
	"fmt"
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/pkg/repository"
	"fr33d0mz/moneyflowx/utils"
	"fr33d0mz/moneyflowx/utils/CustomError"
)

type TransactionService struct {
	repo repository.Repository
}

func NewTransactionService(repo repository.Repository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (t *TransactionService) GetTransactions(userID string, query *dto.TransactionRequestQuery) ([]*models.Transaction, error) {
	return t.repo.Transaction.FindAll(userID, query)
}

// TODO: change it
func (t *TransactionService) TopUp(input *dto.TopUpRequestBody) (*models.Transaction, error) {
	wallet, err := t.repo.Wallet.FindByUserId(input.User.ID)
	if err != nil {
		return &models.Transaction{}, err
	}

	if wallet.ID == "" {
		return &models.Transaction{}, &CustomError.WalletNotFoundError{}
	}

	if wallet.UserType == "unidentified" && wallet.Balance+input.Amount > 10000.0 {
		return &models.Transaction{}, &CustomError.AmountUnidentifiedUserLimitError{}
	} else if wallet.UserType == "identified" && wallet.Balance+input.Amount > 100000.0 {
		return &models.Transaction{}, &CustomError.AmountIdentifiedUserLimitError{}
	}

	_transaction := &models.Transaction{}
	_transaction.UserID = input.User.ID
	_transaction.WalletID = wallet.ID
	_transaction.Amount = input.Amount
	_transaction.Type = "Top up"

	transaction, err := t.repo.Transaction.Save(_transaction)
	if err != nil {
		return &models.Transaction{}, err
	}

	wallet.Balance += input.Amount
	wallet, err = t.repo.Wallet.Update(wallet)
	if err != nil {
		return &models.Transaction{}, err
	}

	transaction.User = *input.User
	transaction.Wallet = *wallet

	return transaction, nil
}

func (t *TransactionService) Transfer(input *dto.TransferRequestBody) (*models.Transaction, error) {
	myWallet, err := t.repo.Wallet.FindByUserId(input.User.ID)
	if err != nil {
		return &models.Transaction{}, err
	}

	if myWallet.ID == "" {
		return &models.Transaction{}, &CustomError.WalletNotFoundError{}
	}

	if !utils.IsWalletNumberValid(input.User.ID, myWallet.Number) {
		return &models.Transaction{}, &CustomError.TransferToSameWalletError{}
	}

	if myWallet.Balance-input.Amount <= -1 {
		return &models.Transaction{}, &CustomError.InsufficientBalanceError{}
	}

	destinationWallet, err := t.repo.Wallet.FindByNumber(input.WalletNumber)
	if err != nil {
		return &models.Transaction{}, err
	}

	if destinationWallet.ID == "" {
		return &models.Transaction{}, &CustomError.WalletNotFoundError{}
	}

	if destinationWallet.UserType == "unidentified" && destinationWallet.Balance+input.Amount > 10000.0 {
		return &models.Transaction{}, &CustomError.AmountUnidentifiedUserLimitError{}
	} else if destinationWallet.UserType == "identified" && destinationWallet.Balance+input.Amount > 100000.0 {
		return &models.Transaction{}, &CustomError.AmountIdentifiedUserLimitError{}
	}

	transaction := &models.Transaction{}
	transaction.UserID = input.User.ID
	transaction.WalletID = destinationWallet.ID
	transaction.Amount = input.Amount
	if input.Comment == "" {
		transaction.Comment = fmt.Sprintf("from %s to %s", input.User.Firstname, destinationWallet.User.Firstname)
	}
	transaction.Comment = input.Comment
	transaction.Type = "Transfer"

	transaction, err = t.repo.Transaction.Save(transaction)
	if err != nil {
		return &models.Transaction{}, err
	}

	myWallet.Balance -= input.Amount
	myWallet, err = t.repo.Wallet.Update(myWallet)
	if err != nil {
		return &models.Transaction{}, err
	}

	destinationWallet.Balance += input.Amount
	destinationWallet, err = t.repo.Wallet.Update(destinationWallet)
	if err != nil {
		return &models.Transaction{}, err
	}

	transaction.User = *input.User
	transaction.Wallet = *destinationWallet

	return transaction, nil
}

func (t *TransactionService) CountTransaction(userID string) (int64, error) {
	return t.repo.Transaction.Count(userID)
}
