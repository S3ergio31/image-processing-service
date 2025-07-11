package domain

type Rotator struct {
}

func (c Rotator) Transform(transformations *Transformations, imageEditor ImageEditor, content []byte) []byte {
	return imageEditor.Rotate(transformations.Rotate, content)
}
