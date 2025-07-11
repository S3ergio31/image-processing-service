package infrastructure

type Resize struct {
	Width  int `json:"width" binding:"required"`
	Height int `json:"height" binding:"required"`
}

type Crop struct {
	Width  int `json:"width" binding:"required"`
	Height int `json:"height" binding:"required"`
	X      int `json:"x" binding:"required"`
	Y      int `json:"y" binding:"required"`
}

type Transformations struct {
	Rotate  *int   `json:"rotate" binding:"required"`
	Format  string `json:"format" binding:"required"`
	*Resize `json:"resize"`
	*Crop   `json:"crop"`
}

type TranformBody struct {
	Transformations `json:"transformations" binding:"required"`
}
