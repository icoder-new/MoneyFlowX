package CustomError

type AmountLimitError struct{}

func (e *AmountLimitError) Error() string {
	return "maximum balance for identified accounts exceeded"
}
