package infrastructure

type Resize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Crop struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

type Filters struct {
	Grayscale bool `json:"grayscale"`
	Sepia     bool `json:"sepia"`
}

type Transformations struct {
	Rotate   int    `json:"rotate"`
	Format   string `json:"format"`
	*Resize  `json:"resize"`
	*Crop    `json:"crop"`
	*Filters `json:"filters"`
}

type TranformBody struct {
	Transformations `json:"transformations"`
}
