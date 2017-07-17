package game

type Map struct {
	Width  int
	Height int
	Grid   [][]int
}

//	constructor
func NewMap(w, h int) *Map {
	grid := make([][]int, w)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, h)
	}
	return &Map{
		Width:  w,
		Height: h,
		Grid:   grid,
	}
}
