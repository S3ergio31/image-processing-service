package domain

type Formater struct {
}

func (c Formater) Transform(transformations *Transformations, imageEditor ImageEditor, content []byte) []byte {
	return imageEditor.Format(transformations.Format, content)
}
