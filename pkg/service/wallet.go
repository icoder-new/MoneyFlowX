package service

import (
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/pkg/repository"
	"fr33d0mz/moneyflowx/utils"
	"fr33d0mz/moneyflowx/utils/CustomError"
)

type WalletService struct {
	repo repository.Repository
}

func NewWalletService(repo repository.Repository) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (w *WalletService) GetWalletByUserId(input *dto.WalletRequestBody) (*models.Wallet, error) {
	wallet, err := w.repo.Wallet.FindByUserId(input.UserID)
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (w *WalletService) CreateWallet(input *dto.WalletRequestBody) (*models.Wallet, error) {
	user, err := w.repo.User.FindById(input.UserID)
	if err != nil {
		return &models.Wallet{}, err
	}

	if user.ID == "" {
		return &models.Wallet{}, &CustomError.UserNotFoundError{}
	}

	wallet, err := w.repo.Wallet.FindByUserId(user.ID)
	if err != nil {
		return &models.Wallet{}, err
	}

	if wallet.ID != "" {
		return &models.Wallet{}, &CustomError.WalletAlreadyExistsError{}
	}

	wallet.UserID = user.ID
	wallet.UserType = user.Type
	wallet.Balance = 0.0
	wallet.IsActive = true
	wallet.Number = utils.GenerateWalletNumber(user.ID)

	return w.repo.Wallet.CreateWallet(wallet)
}
