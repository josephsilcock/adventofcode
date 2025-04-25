package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"strings"
	"testing"
)

func TestGetTotalNumbersAfterNIterations(t *testing.T) {
	var numbers []int
	for _, val := range strings.Fields(utils.ReadFile("input.txt")[0]) {
		numbers = append(numbers, utils.ConvertStringToInt(val))
	}
	total := getTotalNumbersAfterNIterations(numbers, 25)

	if total != 231278 {
		t.Errorf("Got %v want 231278", total)
	}
}

func BenchmarkGetTotalNumbersAfterNIterations(b *testing.B) {
	var numbers []int
	for _, val := range strings.Fields(utils.ReadFile("input.txt")[0]) {
		numbers = append(numbers, utils.ConvertStringToInt(val))
	}

	for b.Loop() {
		getTotalNumbersAfterNIterations(numbers, 25)
	}
}
