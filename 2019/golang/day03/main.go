package main

import (
	"fmt"
	"math"

	"github.com/mpanelo/adventofcode/2019/golang/day03/point"
	"github.com/mpanelo/adventofcode/2019/golang/day03/wireio"
)

func main() {
	// Puzzle 1
	wires := wireio.ScanAllWires("../../puzzledata/day03/input.txt")

	points1 := wires[0].Points()
	points2 := wires[1].Points()

	set1 := points1.Set()
	set2 := points2.Set()

	intersection := set1.Intersection(set2)

	fmt.Println("Puzzle 1:", minManhattanDistance(intersection))
	fmt.Println("Puzzle 2:", minCombinedSteps(intersection))
}

func minManhattanDistance(pairs []point.Pair) float64 {
	minDistance := math.Inf(1)
	for _, pair := range pairs {
		point := pair.First
		minDistance = math.Min(minDistance, point.ManhattanDistance())
	}
	return minDistance
}

func minCombinedSteps(pairs []point.Pair) float64 {
	minSteps := math.Inf(1)
	for _, pair := range pairs {
		totalSteps := pair.First.StepsTaken + pair.Second.StepsTaken
		minSteps = math.Min(minSteps, float64(totalSteps))
	}
	return minSteps
}
