package domain

type UserRepository interface {
	Save(user User)
	UsedUsername(username string) bool
}
