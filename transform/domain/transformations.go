package domain

import "fmt"

type Transformations struct {
	ImageUuid string
	Username  string
	Rotate    int
	Format    string
	*Resize
	*Crop
}

type Transformation struct {
	*Transformations
	Image
	Content []byte
}

func (t Transformation) Path() string {
	return fmt.Sprintf(
		"images/%s/transformations/%s_%s.%s",
		t.Image.Username(),
		t.Uuid(),
		t.Name(),
		t.Format,
	)
}
