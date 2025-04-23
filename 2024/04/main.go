package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
)

func main() {
	rows := utils.ReadFile("input.txt")

	xmasWordSearch := XmasWordSearch{}
	crossMasWordSearch := CrossMasWordSearch{}

	for _, row := range rows {
		xmasWordSearch.AddRow(row)
		crossMasWordSearch.AddRow(row)
	}

	fmt.Println("Star 1:", xmasWordSearch.FindCountOfXmas())
	fmt.Println("Star 2:", crossMasWordSearch.FindCountOfCrossMas())
}
