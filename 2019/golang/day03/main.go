package main

import (
	"fmt"
	"math"

	"github.com/mpanelo/adventofcode/2019/golang/day03/types"
	"github.com/mpanelo/adventofcode/2019/golang/day03/wireio"
)

type Points []types.Point

func main() {
	// Puzzle 1
	wires := wireio.ScanAllWires("../../puzzledata/day03/input.txt")
	var allPoints []Points

	for _, wire := range wires {
		allPoints = append(allPoints, traceWirePath(wire))
	}

	intersection := allPoints[0].intersection(allPoints[1])
	fmt.Println(intersection)
	minManDist := math.Inf(1)

	for _, point := range intersection {
		x := math.Abs(float64(point.X))
		y := math.Abs(float64(point.Y))
		minManDist = math.Min(minManDist, x+y)
	}

	fmt.Println(minManDist)
}

func traceWirePath(wire *types.Wire) Points {
	var x, y int

	points := Points{types.Point{X: x, Y: y}}

	for _, wireSection := range wire.Path {
		switch wireSection.Direction {
		case 'U':
			for i := 0; i < wireSection.Steps; i++ {
				y = y + 1
				points = append(points, types.Point{X: x, Y: y})
			}
		case 'R':
			for i := 0; i < wireSection.Steps; i++ {
				x = x + 1
				points = append(points, types.Point{X: x, Y: y})
			}
		case 'D':
			for i := 0; i < wireSection.Steps; i++ {
				y = y - 1
				points = append(points, types.Point{X: x, Y: y})
			}
		case 'L':
			for i := 0; i < wireSection.Steps; i++ {
				x = x - 1
				points = append(points, types.Point{X: x, Y: y})
			}
		}
	}

	return points
}

func (p Points) intersection(other Points) Points {
	var intersection Points
	seen := make(map[string]bool)

	for _, point := range p {
		if point.X == 0 && point.Y == 0 {
			continue
		}

		key := fmt.Sprintf("%d%d", point.X, point.Y)
		if _, ok := seen[key]; !ok {
			seen[key] = true
		}
	}

	for _, point := range other {
		if point.X == 0 && point.Y == 0 {
			continue
		}

		key := fmt.Sprintf("%d%d", point.X, point.Y)
		if _, ok := seen[key]; ok {
			intersection = append(intersection, point)
			delete(seen, key)
		}
	}

	return intersection
}
