package domain

type Hasher interface {
	CompareHashAndPassword(hashedPassword string, password string) error
}
