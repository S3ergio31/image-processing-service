package domain

type InvalidTokenGeneration struct {
	Err error
}

func (i InvalidTokenGeneration) Error() string {
	return "InvalidTokenGeneration: cannot generate user token: " + i.Err.Error()
}
