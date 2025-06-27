package domain

type UserRepository interface {
	Save(user User)
	FindByUsername(username string) (User, *UserNotFound)
}
