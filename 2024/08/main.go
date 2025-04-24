package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
)

func main() {
	rows := utils.ReadFile("input.txt")
	noResonanceAntinodes, resonanceAntinodes := getAntinodes(rows)

	fmt.Println("Star 1:", noResonanceAntinodes)
	fmt.Println("Star 2:", resonanceAntinodes)
}

func getAntinodes(rows []string) (int, int) {
	noResonanceMap := AntennaeMap{
		antennaeMapping: make(map[rune][]Point),
		height:          len(rows),
		width:           len(rows[0]),
		antinodeMapping: make(map[Point]bool),
		antennaAdder:    NoResonanceAdder{},
	}
	resonanceMap := AntennaeMap{
		antennaeMapping: make(map[rune][]Point),
		height:          len(rows),
		width:           len(rows[0]),
		antinodeMapping: make(map[Point]bool),
		antennaAdder:    ResonanceAdder{},
	}

	for rowIndex, row := range rows {
		noResonanceMap.AddRow(row, rowIndex)
		resonanceMap.AddRow(row, rowIndex)
	}

	return noResonanceMap.GetTotalAntinodes(), resonanceMap.GetTotalAntinodes()
}
