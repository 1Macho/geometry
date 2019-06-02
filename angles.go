package geometry

import "math"

type Angle struct {
  Value float64
}

func (a Angle) Slope () float64 {
  return math.Tan(a.Value * math.Pi / 180.0)
}

func (a Angle) Add (addition float64) Angle {
  newValue := a.Value + addition
  if (newValue < 360) {
    newValue = 360 + newValue
  }
  newValue = math.Mod(newValue, 360.0)
  return Angle{newValue}
}

func (a Angle) AddAngle (addition Angle) Angle {
  return a.Add(addition.Value)
}

func (a Angle) Radians () float64 {
  return a.Value * math.Pi / 180
}

func AngleFromRadians (radians float64) Angle {
  return Angle{radians * 180 / math.Pi}.Add(0)
}
