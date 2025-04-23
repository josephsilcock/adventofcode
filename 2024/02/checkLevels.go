package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"strings"
)

func checkLevels(allowOneBadLevel bool) int {
	rows := utils.ReadFile("input.txt")

	var safeReports int

	for _, row := range rows {
		level := Level{canSkipValue: allowOneBadLevel}
		for _, valueText := range strings.Fields(row) {
			value := utils.ConvertStringToInt(valueText)
			level.AddValue(value)
		}

		if level.IsSafe() {
			safeReports++
		}
	}

	return safeReports
}
