package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"strings"
)

func getRules() map[string][]string {
	rules := make(map[string][]string)

	rows := utils.ReadFile("input_rules.txt")

	for _, row := range rows {
		numbers := strings.Split(row, "|")

		rules[numbers[0]] = append(rules[numbers[0]], numbers[1])
	}

	return rules
}
