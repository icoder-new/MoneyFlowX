package CustomError

type InsufficientBalanceError struct {
}

func (e *InsufficientBalanceError) Error() string {
	return "insufficient balance"
}
