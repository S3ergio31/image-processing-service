package transformations

type Transformations struct {
	Rotate int
	Format string
	*Resize
	*Crop
	*Filters
}
