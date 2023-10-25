package service

import (
	"errors"
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/pkg/repository"
	"fr33d0mz/moneyflowx/utils"
	"github.com/spf13/cast"
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

func (t *TransactionService) TopUp(input *dto.TopUpRequestBody) (*models.Transaction, error) {
	sourceOfFund, err := t.repo.SourceOfFund.FindById(input.SourceOfFundID)
	if err != nil {
		return &models.Transaction{}, err
	}

	if sourceOfFund.ID == "" {
		return &models.Transaction{}, errors.New("`source of fund` not found")
	}

	wallet, err := t.repo.Wallet.FindByUserId(input.User.ID)
	if err != nil {
		return &models.Transaction{}, err
	}

	if wallet.ID == "" {
		return &models.Transaction{}, errors.New("wallet not found error")
	}

	transaction := &models.Transaction{}
	transaction.SourceOfFundID = &sourceOfFund.ID
	transaction.UserID = input.User.ID
	transaction.WalletID = wallet.ID
	transaction.Amount = input.Amount
	transaction.Comment = "Top Up from " + sourceOfFund.Name
	transaction.Type = "Top Up"

	_transaction, err := t.repo.Transaction.Save(transaction)
	if err != nil {
		return transaction, err
	}

	wallet.Balance = wallet.Balance + input.Amount
	wallet, err = t.repo.Wallet.Update(wallet)
	if err != nil {
		return _transaction, err
	}

	transaction.SourceOfFund = sourceOfFund
	transaction.User = *input.User
	transaction.Wallet = *wallet

	return transaction, nil

}

func (t *TransactionService) Transfer(input *dto.TransferRequestBody) (*models.Transaction, error) {
	myWallet, err := t.repo.Wallet.FindByNumber(input.WalletNumber)
	if err != nil {
		return &models.Transaction{}, err
	}

	if myWallet.ID == "" {
		return &models.Transaction{}, errors.New("wallet not found")
	}

	if myWallet.Balance < input.Amount {
		return &models.Transaction{}, errors.New("insufficient balance")
	}

	if utils.IsWalletNumberValid(myWallet.UserID, myWallet.Number) {
		return &models.Transaction{}, errors.New("transfer to the same wallet")
	}

	destinationWallet, err := t.repo.Wallet.FindByNumber(myWallet.Number)
	if err != nil {
		return &models.Transaction{}, err
	}

	if destinationWallet.ID == "" {
		return &models.Transaction{}, errors.New("wallet not found")
	}

	transaction := &models.Transaction{}
	transaction.UserID = input.User.ID
	transaction.WalletID = destinationWallet.ID
	transaction.Amount = input.Amount
	transaction.Comment = input.Comment
	transaction.Type = "Send money"

	transaction, err = t.repo.Transaction.Save(transaction)
	if err != nil {
		return &models.Transaction{}, err
	}

	myWallet.Balance = myWallet.Balance - input.Amount
	myWallet, err = t.repo.Wallet.Update(myWallet)
	if err != nil {
		return &models.Transaction{}, err
	}

	destinationWallet.Balance = destinationWallet.Balance + input.Amount
	_, err = t.repo.Wallet.Update(destinationWallet)
	if err != nil {
		return &models.Transaction{}, err
	}

	balance := cast.ToString(myWallet.Balance)
	transaction.SourceOfFundID = &balance
	transaction.User = *input.User
	transaction.Wallet = *destinationWallet
	return transaction, nil
}

func (t *TransactionService) CountTransaction(userID string) (int64, error) {
	return t.repo.Transaction.Count(userID)
}
