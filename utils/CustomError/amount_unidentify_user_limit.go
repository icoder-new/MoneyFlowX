package CustomError

type AmountUnidentifiedUserLimitError struct{}

func (e *AmountUnidentifiedUserLimitError) Error() string {
	return "the maximum balance for unidentified accounts has been exceeded"
}
