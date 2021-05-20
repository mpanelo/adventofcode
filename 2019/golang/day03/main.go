package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mpanelo/adventofcode/2019/golang/day03/mathutil"
	"github.com/mpanelo/adventofcode/2019/golang/day03/types"
	"github.com/mpanelo/adventofcode/2019/golang/day03/wireio"
)

const (
	FWIRE = iota
	SWIRE
	CROSS
	EMPTY
)

func main() {
	// Puzzle 1
	firstWire, secondWire := wireio.ScanWires("../../puzzledata/day03/input.txt")

	fBounds := getGridBoundary(firstWire)
	sBounds := getGridBoundary(secondWire)

	fmt.Println(fBounds)
	fmt.Println(sBounds)
}

func getGridBoundary(wirePath []string) types.GridBoundary {
	var minX, minY, maxX, maxY int
	var x, y int

	for _, subpath := range wirePath {
		runes := []rune(subpath)
		direction := runes[0]
		steps, err := strconv.Atoi(string(runes[1:]))
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 'L':
			x = x - steps
			minX = mathutil.Min(minX, x)
		case 'R':
			x = x + steps
			maxX = mathutil.Max(maxX, x)
		case 'D':
			y = y - steps
			minY = mathutil.Min(minY, y)
		case 'U':
			y = y + steps
			maxY = mathutil.Max(maxY, y)
		}
	}

	return types.GridBoundary{
		MinX: minX,
		MaxX: maxX,
		MinY: minY,
		MaxY: maxY,
	}
}