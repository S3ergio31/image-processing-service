package domain

type TransformationChain struct {
	Transformers []Transformer
	ImageEditor  ImageEditor
}

func (t *TransformationChain) Transform(transformations *Transformations, content []byte) []byte {
	result := content

	for _, transformer := range t.Transformers {
		content = transformer.Transform(transformations, t.ImageEditor, result)
		if content == nil {
			continue
		}
		result = content
	}

	return result
}
