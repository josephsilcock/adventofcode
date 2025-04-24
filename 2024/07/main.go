package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
)

func main() {
	rows := utils.ReadFile("input.txt")
	var totalCalibration int
	var totalCalibrationIncludingCombination int

	for _, row := range rows {
		target, values := parseInputLine(row)
		if isCalculationPossible(target, values) {
			totalCalibration += target
		}
		if isCalculationPossibleIncludingCombination(target, values) {
			totalCalibrationIncludingCombination += target
		}
	}

	fmt.Println("Star 1:", totalCalibration)
	fmt.Println("Star 1:", totalCalibrationIncludingCombination)
}
