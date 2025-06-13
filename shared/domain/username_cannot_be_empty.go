package domain

type UsernameCannotBeEmpty struct {
}

func (UsernameCannotBeEmpty) Error() string {
	return "username cannot be empty"
}
