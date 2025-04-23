package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"regexp"
	"strconv"
)

func getBasicTotal() (total int) {
	rows := utils.ReadFile("input.txt")

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, row := range rows {
		matches := regex.FindAllStringSubmatch(row, -1)
		for _, match := range matches {
			number1, _ := strconv.Atoi(match[1])
			number2, _ := strconv.Atoi(match[2])
			total += number1 * number2
		}
	}

	return
}
