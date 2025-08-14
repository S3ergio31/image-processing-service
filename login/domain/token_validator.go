package domain

type TokenValidator interface {
	Username(secret string, token string) string
	IsValid(secret string, token string) bool
}
