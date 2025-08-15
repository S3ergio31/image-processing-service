package domain

type Hasher interface {
	Hash(value string) (string, error)
}
