package domain

type InvalidUuid struct {
}

func (InvalidUuid) Error() string {
	return "invalid v4 uuid"
}
