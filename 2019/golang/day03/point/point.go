package point

import (
	"fmt"
	"math"
)

type Points []Point
type Set map[string]Point

type Point struct {
	X          int
	Y          int
	StepsTaken int
}

type Pair struct {
	First  Point
	Second Point
}

func (p Points) Set() Set {
	set := make(Set)

	for _, point := range p {
		key := point.String()

		if seenPoint, ok := set[key]; !ok {
			set[key] = point
		} else {
			if point.StepsTaken < seenPoint.StepsTaken {
				set[key] = point
			}
		}
	}

	return set
}

func (s Set) Intersection(otherSet Set) []Pair {
	var intersection []Pair

	for k, v := range s {
		if otherV, ok := otherSet[k]; ok {
			intersection = append(intersection, Pair{v, otherV})
		}
	}

	return intersection
}
func (p Point) ManhattanDistance() float64 {
	x := math.Abs(float64(p.X))
	y := math.Abs(float64(p.Y))

	return x + y
}

func (p Point) String() string {
	return fmt.Sprintf("%d%d", p.X, p.Y)
}
