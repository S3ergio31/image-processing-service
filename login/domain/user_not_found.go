package domain

type UserNotFound struct {
}

func (UserNotFound) Error() string {
	return "user not found"
}
