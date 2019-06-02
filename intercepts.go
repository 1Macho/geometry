package geometry

import "math"

func InterceptLineLine (this Line, other Line) Point {
  m := this.Direction.Slope()
  n := other.Direction.Slope()
  if (m == n) {
    return Point{math.Inf(1), math.Inf(1)}
  }
  b := this.TrueOrigin().Y
  c := other.TrueOrigin().Y
  interceptX := (b-c)/(n-m)
  interceptY := (m*c-n*b)/(m-n)
  return Point{interceptX, interceptY}
}

func almostEqual(a, b float64) bool {
  //ai, bi := int64(math.Float64bits(a)), int64(math.Float64bits(b))
  //return a == b || ai-bi <= 1 && -1 <= ai-bi
  //println(math.Abs(a - b))
  return math.Abs(a - b) < 0.001
}

func InterceptRayLine (this Ray, other Line) Point {
  lineIntercept := InterceptLineLine(this.Line, other)
  interceptDirection := this.Line.Origin.Direction(lineIntercept)
  if (almostEqual(interceptDirection.Value, this.Direction.Value)) {
  //if (true) {
    return lineIntercept
  }
  return Point{math.Inf(1), math.Inf(1)}
}

func InterceptRayRay (this Ray, other Ray) Point {
  interceptOne := InterceptRayLine(this, other.Line)
  interceptTwo := InterceptRayLine(other, this.Line)
  //println(interceptTwo.X)
  //println(interceptTwo.Y)
  if (interceptOne.X == math.Inf(1) || interceptOne.Y == math.Inf(1) || interceptTwo.X == math.Inf(1) || interceptTwo.Y == math.Inf(1)) {
    return Point{math.Inf(1), math.Inf(1)}
  }
  return interceptOne
}

func InterceptRaySegment (this Ray, other Segment) Point {
  otherRay := Ray{other.Line}
  intercept := InterceptRayRay(this, otherRay)
  //println(intercept.X)
  //println(intercept.Y)
  if (intercept.X == math.Inf(1) || intercept.Y == math.Inf(1)) {
    return Point{math.Inf(1), math.Inf(1)}
  }
  if (other.Line.Origin.Distance(intercept) <= other.Length || almostEqual(other.Line.Origin.Distance(intercept), other.Length)) {
  //if (true) {
    //println(intercept.X)
    //println(intercept.Y)
    return intercept
  }
  return Point{math.Inf(1), math.Inf(1)}
}
