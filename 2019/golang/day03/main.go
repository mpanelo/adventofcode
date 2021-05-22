package main

import (
	"fmt"
	"math"

	"github.com/mpanelo/adventofcode/2019/golang/day03/types"
	"github.com/mpanelo/adventofcode/2019/golang/day03/wireio"
)

type Pair struct {
	First  types.Point
	Second types.Point
}

type Points []types.Point
type PointSet map[string]types.Point

func main() {
	// Puzzle 1
	wires := wireio.ScanAllWires("../../puzzledata/day03/input.txt")

	firstWirePoints := traceWirePath(wires[0])
	secondWirePoints := traceWirePath(wires[1])

	fSet := firstWirePoints.toPointSet()
	sSet := secondWirePoints.toPointSet()

	intersection := fSet.intersection(sSet)

	fmt.Println("Puzzle 1:", minManhattanDistance(intersection))
	fmt.Println("Puzzle 2:", minCombinedSteps(intersection))
}

func minManhattanDistance(pairs []Pair) float64 {
	minDistance := math.Inf(1)
	for _, pair := range pairs {
		point := pair.First
		minDistance = math.Min(minDistance, point.ManhattanDistance())
	}
	return minDistance
}

func minCombinedSteps(pairs []Pair) float64 {
	minSteps := math.Inf(1)
	for _, pair := range pairs {
		totalSteps := pair.First.StepsTaken + pair.Second.StepsTaken
		minSteps = math.Min(minSteps, float64(totalSteps))
	}
	return minSteps
}

func traceWirePath(wire *types.Wire) Points {
	var x, y, stepsTaken int
	var points Points

	for _, wireSection := range wire.Path {
		for i := 0; i < wireSection.Steps; i++ {
			switch wireSection.Direction {
			case 'U':
				y = y + 1
			case 'R':
				x = x + 1
			case 'D':
				y = y - 1
			case 'L':
				x = x - 1
			}

			stepsTaken++
			points = append(points, types.Point{X: x, Y: y, StepsTaken: stepsTaken})
		}
	}

	return points
}

func (p Points) toPointSet() PointSet {
	set := make(PointSet)

	for _, point := range p {
		key := fmt.Sprintf("%d%d", point.X, point.Y)

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

func (ps PointSet) intersection(otherPs PointSet) []Pair {
	var intersection []Pair

	for k, v := range ps {
		if otherV, ok := otherPs[k]; ok {
			intersection = append(intersection, Pair{v, otherV})
		}
	}

	return intersection
}
