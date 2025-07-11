package domain

type Transformer interface {
	Transform(*Transformations, ImageEditor, []byte) []byte
}
