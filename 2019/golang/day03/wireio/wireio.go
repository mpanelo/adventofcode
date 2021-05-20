package wireio

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ScanWires(datapath string) ([]string, []string) {
	file, err := os.Open(datapath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	firstWire := scanWire(scanner)
	secondWire := scanWire(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return firstWire, secondWire
}

func scanWire(scanner *bufio.Scanner) []string {
	var wire []string
	if ok := scanner.Scan(); ok {
		line := scanner.Text()
		wire = strings.Split(strings.Trim(line, "\n"), ",")
	}
	return wire
}
