package domain

type Croper struct {
}

func (c Croper) Transform(transformations *Transformations, imageEditor ImageEditor, content []byte) []byte {
	if transformations.Crop == nil {
		return nil
	}

	return imageEditor.Crop(transformations.Crop, content)
}
