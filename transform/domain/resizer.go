package domain

type Resizer struct {
}

func (c Resizer) Transform(transformations *Transformations, imageEditor ImageEditor, content []byte) []byte {
	if transformations.Resize == nil {
		return nil
	}

	return imageEditor.Resize(transformations.Resize, content)
}
