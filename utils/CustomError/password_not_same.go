package CustomError

type PasswordNotSameError struct{}

func (e *PasswordNotSameError) Error() string {
	return "password is not the same as confirm password"
}
