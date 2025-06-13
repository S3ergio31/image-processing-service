package infrastructure

type RegisterBody struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
