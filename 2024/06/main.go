package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	initialMaze := createMaze()

	mazeToRun := initialMaze.CreateCopy()
	runMazeAndReturnInLoop(mazeToRun)

	fmt.Println("Star 1:", mazeToRun.visitedPositions)
	fmt.Println("Star 2:", findNumberOfMazesWithLoops(initialMaze, mazeToRun))
}

func createMaze() (maze *Maze) {
	maze = &Maze{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze.AddRow(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func runMazeAndReturnInLoop(maze *Maze) bool {
	maze.StartMaze()
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
				newMaze := maze.CreateCopy()
				newMaze.GetPositionAtPoint(Point{rowIndex, columnIndex}).obstacle = true
				if runMazeAndReturnInLoop(newMaze) {
					total++
				}
			}
		}
	}
	return
}
