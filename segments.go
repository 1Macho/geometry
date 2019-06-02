package geometry

type Segment struct {
  Length float64
  Line
}

func (this *Segment) EndPoint () Point {
  return this.Line.Origin.Translate(this.Length,this.Line.Direction)
}

func SegmentFromPoints (a Point, b Point) Segment {
  newDirection := a.Direction(b)
  newLength := a.Distance(b)
  newLine := Line{newDirection, a}
  return Segment{newLength, newLine}
}

func ShapeFromPoints (points []Point) []Segment {
  result := make([]Segment, len(points))
  if (len(points) >= 3) {
    for index, point := range points {
      if (index == len(points) - 1) {
        result[index] = SegmentFromPoints(point, points[0])
      } else {
        result[index] = SegmentFromPoints(point, points[index + 1])
      }
    }
  }
  return result
}

func IntermitentShapeFromPoints (points []Point) []Segment {
  result := make([]Segment, len(points)/2)
  if (len(points) >= 2 && len(points) % 2 == 0) {
    for index, point := range points {
      if (index % 2 == 0) {
        result[index/2] = SegmentFromPoints(point, points[index + 1])
      }
    }
  }
  return result
}
