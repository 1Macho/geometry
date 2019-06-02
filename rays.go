package geometry

type Ray struct {
  Line
}

func RayFromPoints (a Point, b Point) Ray {
  newDirection := a.Direction(b)
  newLine := Line{newDirection, a}
  return Ray{newLine}
}
