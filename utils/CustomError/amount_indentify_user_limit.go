package CustomError

type AmountIdentifiedUserLimitError struct{}

func (e *AmountIdentifiedUserLimitError) Error() string {
	return "maximum balance for identified accounts exceeded"
}
