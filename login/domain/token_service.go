package domain

type TokenService interface {
	Generate(secret string, username string) (string, error)
}
