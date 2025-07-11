package domain

type ImageEditor interface {
	Resize(*Resize, []byte) []byte
	Crop(*Crop, []byte) []byte
	Rotate(int, []byte) []byte
	Format(string, []byte) []byte
}
