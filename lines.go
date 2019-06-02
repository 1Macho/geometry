package geometry

type Line struct {
  Direction Angle
  Origin Point
}

func (l Line) TrueOrigin () Point {
  if (l.Origin.X == 0) {
    return l.Origin
  }
  deltaX := 0 - l.Origin.X
  deltaY := l.Direction.Slope() * deltaX
  finalY := l.Origin.Y + deltaY
  return Point{0, finalY}
}
