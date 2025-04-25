package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
)

var (
	directions = [4]Point{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
)

type TopographicalMap struct {
	mapping [][]int
}

func (tm *TopographicalMap) AddRow(row string) {
	runes := []rune(row)
	heights := make([]int, len(runes))
	for i, file := range runes {
		heights[i] = utils.ConvertStringToInt(string(file))
	}
	tm.mapping = append(tm.mapping, heights)
}

func (tm *TopographicalMap) pointInBounds(point Point) bool {
	return point.x >= 0 && point.x < len(tm.mapping[0]) && point.y >= 0 && point.y < len(tm.mapping)
}

func (tm *TopographicalMap) getHeightAtPoint(point Point) int {
	return tm.mapping[point.y][point.x]
}

func (tm *TopographicalMap) getUphillPoints(point Point) (uphillPoints []Point) {
	pointHeight := tm.getHeightAtPoint(point)
	for _, direction := range directions {
		newPoint := point.Add(direction)
		if tm.pointInBounds(newPoint) && tm.getHeightAtPoint(newPoint)-pointHeight == 1 {
			uphillPoints = append(uphillPoints, newPoint)
		}
	}
	return
}

func (tm *TopographicalMap) getTrailheadScore(trailheadPoint Point, includeDuplicates bool) int {
	peaksFound := make(map[Point]bool)
	totalTrailsFound := 0

	pointsToSearch := []Point{trailheadPoint}
	for len(pointsToSearch) > 0 {
		point := pointsToSearch[0]
		if tm.getHeightAtPoint(point) == 9 {
			totalTrailsFound++
			peaksFound[point] = true
		} else {
			pointsToSearch = append(pointsToSearch, tm.getUphillPoints(point)...)
		}
		pointsToSearch = pointsToSearch[1:]
	}

	if includeDuplicates {
		return totalTrailsFound
	}
	return len(peaksFound)
}

func (tm *TopographicalMap) GetTotalTrailheadScore(includeDuplicates bool) (total int) {
	for rowIndex, row := range tm.mapping {
		for columnIndex, columnHeight := range row {
			if columnHeight == 0 {
				total += tm.getTrailheadScore(Point{columnIndex, rowIndex}, includeDuplicates)
			}
		}
	}
	return
}
