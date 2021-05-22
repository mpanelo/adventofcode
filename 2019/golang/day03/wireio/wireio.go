package wireio

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mpanelo/adventofcode/2019/golang/day03/point"
)

type Wire struct {
	Path []WireSection
}

type WireSection struct {
	Direction rune
	Steps     int
}

func (w Wire) Points() point.Points {
	var x, y, stepsTaken int
	var points point.Points

	for _, wireSection := range w.Path {
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
			points = append(points, point.Point{X: x, Y: y, StepsTaken: stepsTaken})
		}
	}

	return points

}

func ScanAllWires(datapath string) []*Wire {
	file, err := os.Open(datapath)
	if err != nil {
		log.Fatal(err)
	}

	var wires []*Wire

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rawWire := strings.Split(strings.Trim(line, "\n"), ",")
		wires = append(wires, newWire(rawWire))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wires
}

func newWire(rawWire []string) *Wire {
	wire := Wire{}

	for _, rawSection := range rawWire {
		runes := []rune(rawSection)

		wireSection := WireSection{}

		wireSection.Direction = runes[0]
		steps, err := strconv.Atoi(string(runes[1:]))
		if err != nil {
			log.Fatal(err)
		}
		wireSection.Steps = steps

		wire.Path = append(wire.Path, wireSection)
	}

	return &wire
}
