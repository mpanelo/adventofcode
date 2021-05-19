package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	ADD  = 1
	MULT = 2
	HALT = 99
)

type Program struct {
	ptr       int
	mem       []int
	staticmem []int
}

func NewProgram(datapath string) *Program {
	rawContent, err := ioutil.ReadFile(datapath)
	if err != nil {
		log.Fatal(err)
	}
	content := strings.Trim(string(rawContent), "\n")
	strintegers := strings.Split(content, ",")

	program := Program{}

	for _, strint := range strintegers {
		integer, err := strconv.Atoi(strint)
		if err != nil {
			log.Fatal(err)
		}
		program.staticmem = append(program.staticmem, integer)
	}

	program.mem = make([]int, len(program.staticmem))
	program.Reset()
	return &program
}

func (p *Program) Run() {
	for {
		opcode := p.MovePtr()

		switch opcode {
		case ADD:
			addr1 := p.MovePtr()
			addr2 := p.MovePtr()
			dest := p.MovePtr()

			p.Write(dest, p.Read(addr1)+p.Read(addr2))
		case MULT:
			addr1 := p.MovePtr()
			addr2 := p.MovePtr()
			dest := p.MovePtr()

			p.Write(dest, p.Read(addr1)*p.Read(addr2))
		case HALT:
			return
		}
	}
}

func (p *Program) MovePtr() int {
	read := p.mem[p.ptr]
	p.ptr += 1
	return read
}

func (p *Program) Read(addr int) int {
	return p.mem[addr]
}

func (p *Program) Write(addr int, value int) {
	p.mem[addr] = value
}

func (p *Program) Reset() {
	copy(p.mem, p.staticmem)
	p.ptr = 0
}

func main() {
	program := NewProgram("../../puzzledata/day02/input.txt")

	// Puzzle 1
	program.Write(1, 12)
	program.Write(2, 2)
	program.Run()
	fmt.Println("Position 0:", program.Read(0))
	program.Reset()

	// Puzzle 2
	target := 19690720

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// noun
			program.Write(1, i)
			// verb
			program.Write(2, j)

			program.Run()
			if program.Read(0) == target {
				fmt.Println("100 * noun + verb =", 100*i+j)
				return
			}
			program.Reset()
		}
	}
}
