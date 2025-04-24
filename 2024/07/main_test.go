package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"testing"
)

func TestFindNumberOfMazesWithLoops(t *testing.T) {
	rows := utils.ReadFile("input.txt")
	totalCalibration, totalCalibrationIncludingCombination := processRows(rows)

	if totalCalibration != 5512534574980 {
		t.Errorf("Star1: Got %v want 5512534574980", totalCalibration)
	}

	if totalCalibrationIncludingCombination != 328790210468594 {
		t.Errorf("Star1: Got %v want 328790210468594", totalCalibrationIncludingCombination)
	}
}

func BenchmarkProcessRows(b *testing.B) {
	rows := utils.ReadFile("input.txt")

	for b.Loop() {
		processRows(rows)
	}
}
