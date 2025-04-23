package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
)

func main() {
	initialMaze := createMaze()

	mazeToRun := initialMaze.CreateCopy()
	runMazeAndReturnInLoop(mazeToRun, FacingUp)

	fmt.Println("Star 1:", mazeToRun.visitedPositions)
	fmt.Println("Star 2:", findNumberOfMazesWithLoops(initialMaze, mazeToRun))
}

func createMaze() (maze *Maze) {
	maze = &Maze{}
	rows := utils.ReadFile("input.txt")

	for _, row := range rows {
		maze.AddRow(row)
	}

	return
}

func runMazeAndReturnInLoop(maze *Maze, direction FacingDirection) bool {
	maze.StartMaze(direction)
	for {
		if maze.WillLeaveBounds() {
			return false
		}
		if maze.WillHaveBeenHereBefore() {
			return true
		}
		maze.Move()
	}
}

func findNumberOfMazesWithLoops(maze *Maze, runMaze *Maze) (total int) {
	for columnIndex, column := range runMaze.positions {
		for rowIndex, position := range column {
			if len(position.visitedDirections) != 0 {
				firstVisitedDirection := position.visitedDirections[0]
				obstaclePoint := Point{rowIndex, columnIndex}
				newMaze := maze.CreateCopy()
				newMaze.GetPositionAtPoint(Point{rowIndex, columnIndex}).obstacle = true
				newMaze.startingLocation = obstaclePoint.Add(Point(firstVisitedDirection).Scale(-1))
				if runMazeAndReturnInLoop(newMaze, firstVisitedDirection) {
					total++
				}
			}
		}
	}
	return
}
