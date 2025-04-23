package main

import (
	"testing"
)

func TestFindNumberOfMazesWithLoops(t *testing.T) {
	initialMaze := createMaze()

	mazeToRun := initialMaze.CreateCopy()
	runMazeAndReturnInLoop(mazeToRun, FacingUp)

	star1 := mazeToRun.visitedPositions

	if star1 != 5516 {
		t.Errorf("Star1: Got %v want 5516", star1)
	}

	star2 := findNumberOfMazesWithLoops(initialMaze, mazeToRun)

	if star2 != 2008 {
		t.Errorf("Star2: Got %v want 2008", star2)
	}

}

func BenchmarkFindNumberOfMazesWithLoops(b *testing.B) {
	initialMaze := createMaze()

	mazeToRun := initialMaze.CreateCopy()
	runMazeAndReturnInLoop(mazeToRun, FacingUp)

	for b.Loop() {
		findNumberOfMazesWithLoops(initialMaze, mazeToRun)
	}
}
