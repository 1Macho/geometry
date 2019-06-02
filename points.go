package geometry

import "math"

type Point struct {
  X float64
  Y float64
}

func (this Point) Direction (other Point) Angle {
  dx := other.X - this.X
  dy := other.Y - this.Y
  radians := math.Atan2(dy, dx)
  return AngleFromRadians(radians)
}

func (this Point) Distance (other Point) float64 {
  dx := other.X - this.X
  dy := other.Y - this.Y
  distanceSquared := (dx * dx) + (dy * dy)
  return math.Sqrt(distanceSquared)
}

func (this Point) Translate (distance float64, direction Angle) Point {
  xChange := math.Cos(direction.Radians()) * distance
  yChange := math.Sin(direction.Radians()) * distance
  return Point{this.X + xChange, this.Y + yChange}
}

func (this Point) Inverse () Point {
  return Point{this.X * -1, this.Y * -1}
}

func (this Point) Multiply (value float64) Point {
  return Point{this.X * value, this.Y * value}
}

func (this Point) Add (other Point) Point {
  return Point{this.X+other.X,this.Y+other.Y}
}

func (this Point) Heading () Angle {
  return Point{0,0}.Direction(this)
}

func (this Point) Magnitude () float64 {
  return this.Distance(Point{0,0})
}
