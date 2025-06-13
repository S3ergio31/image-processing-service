package domain

type UserAlreadyExists struct {
}

func (UserAlreadyExists) Error() string {
	return "user is already registered"
}
