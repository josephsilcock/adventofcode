package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"strings"
)

func anyForbiddenItems[T comparable](itemsToCheck, forbiddenItems []T) (bool, int) {
	lookup := make(map[T]bool)
	for _, item := range forbiddenItems {
		lookup[item] = true
	}

	for index, item := range itemsToCheck {
		if lookup[item] {
			return true, index
		}
	}

	return false, 0
}

func isUpdateOrdered(update []string, rules map[string][]string) (bool, int, int) {
	var updateSoFar []string
	for index, page := range update {
		forbiddenItems, foundIndex := anyForbiddenItems(updateSoFar, rules[page])
		if forbiddenItems {
			return false, index, foundIndex
		}
		updateSoFar = append(updateSoFar, page)
	}
	return true, 0, 0
}

func categoriseUpdates(rules map[string][]string) (ordered, unordered [][]string) {
	rows := utils.ReadFile("input_updates.txt")

	for _, row := range rows {
		update := strings.Split(row, ",")

		if isOrdered, _, _ := isUpdateOrdered(update, rules); isOrdered {
			ordered = append(ordered, update)
		} else {
			unordered = append(unordered, update)
		}
	}
	return
}
