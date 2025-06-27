package domain

type InvalidCredentials struct {
}

func (InvalidCredentials) Error() string {
	return "incorrect username or password"
}
