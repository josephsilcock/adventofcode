package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"strings"
)

func parseInputLine(input string) (int, []int) {
	parts := strings.Split(input, ":")

	target := utils.ConvertStringToInt(parts[0])

	var values []int
	for _, value := range strings.Fields(parts[1]) {
		values = append(values, utils.ConvertStringToInt(value))
	}

	return target, values
}
