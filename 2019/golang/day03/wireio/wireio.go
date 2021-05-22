package wireio

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mpanelo/adventofcode/2019/golang/day03/types"
)

func ScanAllWires(datapath string) []*types.Wire {
	file, err := os.Open(datapath)
	if err != nil {
		log.Fatal(err)
	}

	var wires []*types.Wire

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

func newWire(rawWire []string) *types.Wire {
	wire := types.Wire{}

	for _, rawSection := range rawWire {
		runes := []rune(rawSection)

		wireSection := types.WireSection{}

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
