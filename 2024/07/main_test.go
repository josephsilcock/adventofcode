package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"testing"
)

func TestFindNumberOfMazesWithLoops(t *testing.T) {
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

	if totalCalibration != 5512534574980 {
		t.Errorf("Star1: Got %v want 5512534574980", totalCalibration)
	}

	if totalCalibrationIncludingCombination != 328790210468594 {
		t.Errorf("Star1: Got %v want 328790210468594", totalCalibrationIncludingCombination)
	}
}

func BenchmarkIsCalculationPossible(b *testing.B) {
	rows := utils.ReadFile("input.txt")

	for b.Loop() {
		for _, row := range rows {
			target, values := parseInputLine(row)
			isCalculationPossible(target, values)
		}
	}
}

func BenchmarkIsCalculationPossibleIncludingCombination(b *testing.B) {
	rows := utils.ReadFile("input.txt")

	for b.Loop() {
		for _, row := range rows {
			target, values := parseInputLine(row)
			isCalculationPossibleIncludingCombination(target, values)
		}
	}
}
