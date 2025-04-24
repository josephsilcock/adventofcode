package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"testing"
)

func TestFindNumberOfMazesWithLoops(t *testing.T) {
	rows := utils.ReadFile("input.txt")
	noResonanceAntinodes, resonanceAntinodes := getAntinodes(rows)

	if noResonanceAntinodes != 256 {
		t.Errorf("Star1: Got %v want 256", noResonanceAntinodes)
	}

	if resonanceAntinodes != 1005 {
		t.Errorf("Star1: Got %v want 1005", resonanceAntinodes)
	}
}

func BenchmarkGetAntinodes(b *testing.B) {
	rows := utils.ReadFile("input.txt")

	for b.Loop() {
		getAntinodes(rows)
	}
}
