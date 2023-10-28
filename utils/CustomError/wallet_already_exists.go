package CustomError

type WalletAlreadyExistsError struct{}

func (e *WalletAlreadyExistsError) Error() string {
	return "wallet already exists"
}
