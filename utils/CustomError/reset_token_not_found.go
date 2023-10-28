package CustomError

type ResetTokenNotFoundError struct{}

func (e *ResetTokenNotFoundError) Error() string {
	return "invalid reset token"
}
