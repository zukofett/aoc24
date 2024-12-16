package aoc

type Coordinate struct {
	Position  Point
	Direction Point
}

func (c *Coordinate) TurnRight() {
	switch c.Direction {
	case Down:
		c.Direction = Left
	case Left:
		c.Direction = Up
	case Up:
		c.Direction = Right
	case Right:
		c.Direction = Down
	}
}

func (c *Coordinate) Walk() {
    c.Position.X += c.Direction.X
    c.Position.Y += c.Direction.Y
}

func (c Coordinate) GetNext() Point {
    return Point{
        X: c.Position.X + c.Direction.X,
        Y: c.Position.Y + c.Direction.Y,
    }
}
