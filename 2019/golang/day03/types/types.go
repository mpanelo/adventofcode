package types

import "math"

type Wire struct {
	Path []WireSection
}

type WireSection struct {
	Direction rune
	Steps     int
}

type Point struct {
	X          int
	Y          int
	StepsTaken int
}


func (p Point) ManhattanDistance() float64 {
	x := math.Abs(float64(p.X))
	y := math.Abs(float64(p.Y))

	return x + y
}
