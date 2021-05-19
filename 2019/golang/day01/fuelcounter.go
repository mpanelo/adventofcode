package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Module struct {
	mass int
}

func main() {
	modules := Modules("../../puzzledata/day01/input.txt")
	fmt.Println("Puzzle One Solution:", PuzzleOne(modules))
	fmt.Println("Puzzle Two Solution:", PuzzleTwo(modules))
}

func PuzzleOne(modules []Module) int {
	var sum int
	for _, module := range modules {
		sum += CalculateFuel(module.mass)
	}
	return sum
}

func PuzzleTwo(modules []Module) int {
	var sum int
	for _, module := range modules {
		sum += CalculateFuelRecursive(module.mass, 0)
	}
	return sum
}

func Modules(datapath string) []Module {
	file, err := os.Open(datapath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var modules []Module
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")

		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		modules = append(modules, Module{mass})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return modules
}

func CalculateFuel(mass int) int {
	return mass/3 - 2
}

func CalculateFuelRecursive(mass int, fuelSum int) int {
	fuel := CalculateFuel(mass)
	if fuel <= 0 {
		return fuelSum
	}
	return CalculateFuelRecursive(fuel, fuel+fuelSum)
}
