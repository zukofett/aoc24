package aoc

type Point struct {
	X int
	Y int
}

func (p *Point) AddPoint(other Point) {
	p.X += other.X
	p.Y += other.Y
}

func (p Point) FindDistance(other Point) Point {
	return Point{
		X: other.X - p.X,
		Y: other.Y - p.Y,
	}
}

func (p Point) JoinPoints(other Point) Point {
	return Point{
		X: other.X + p.X,
		Y: other.Y + p.Y,
	}
}

func (p Point) ReversePoint() Point {
	return Point{X: -p.X, Y: -p.Y}
}

var (
	Up    = Point{X: 0, Y: -1}
	Right = Point{X: 1, Y: 0}
	Down  = Point{X: 0, Y: 1}
	Left  = Point{X: -1, Y: 0}
)

var CardinalDirections = []Point{Up, Right, Down, Left}

func (p Point) IsOnGrid(x, y int) bool {
	return IsInRange(p.X, 0, x) && IsInRange(p.Y, 0, y)
}

func OnGrid(y int, x int) func(Point) bool {
	return func(p Point) bool {
		return p.IsOnGrid(x, y)
	}
}
