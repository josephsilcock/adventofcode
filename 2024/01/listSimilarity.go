package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
)

func getListSimilarity() int {
	rows := utils.ReadFile("input.txt")

	list1 := make(map[int]int)
	list2 := make(map[int]int)

	for _, row := range rows {
		numbers := getLineValues(row)
		list1[numbers[0]]++
		list2[numbers[1]]++
	}

	var similarityScore int

	for number, occurrences := range list1 {
		similarityScore += number * occurrences * list2[number]
	}

	return similarityScore
}
