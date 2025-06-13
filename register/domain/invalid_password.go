package domain

type InvalidPassword struct {
}

func (InvalidPassword) Error() string {
	return "invalid password"
}
