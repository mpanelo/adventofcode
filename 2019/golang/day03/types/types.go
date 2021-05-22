package types

type Wire struct {
	Path []WireSection
}

type WireSection struct {
	Direction rune
	Steps     int
}

type Point struct {
	X int
	Y int
}
