package game

type Level struct {
	StartingPos           *Point
	Floor, Walls, Ceiling *Map
}

//	returns the width and height
func (l *Level) GetSize() (int, int) {
	return l.Floor.GetSize()
}

//	returns a boolean indicating if the given position is walkable
func (l *Level) Walkable(p *Point) bool {
	return l.Walls.GetValueAtPoint(p) == 0
}

func NewLevel(sp *Point, f, w, c *Map) *Level {
	return &Level{
		StartingPos: sp,
		Floor:       f,
		Walls:       w,
		Ceiling:     c,
	}
}
