package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
)

func main() {
	rows := utils.ReadFile("input.txt")
	tm := TopographicalMap{}

	for _, row := range rows {
		tm.AddRow(row)
	}

	fmt.Println("Star 1:", tm.GetTotalTrailheadScore(false))
	fmt.Println("Star 2:", tm.GetTotalTrailheadScore(true))
}
