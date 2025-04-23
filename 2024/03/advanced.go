package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"regexp"
	"strconv"
)

func getAdvancedTotal() (total int) {
	rows := utils.ReadFile("input.txt")

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	enabled := true

	for _, row := range rows {
		matches := regex.FindAllStringSubmatch(row, -1)
		for _, match := range matches {
			switch match[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					number1, _ := strconv.Atoi(match[1])
					number2, _ := strconv.Atoi(match[2])
					total += number1 * number2
				}
			}
		}
	}

	return
}
