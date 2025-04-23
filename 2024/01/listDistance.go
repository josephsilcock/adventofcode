package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"math"
	"sort"
)

func getListDistance() int {
	rows := utils.ReadFile("input.txt")

	var list1, list2 []int

	for _, row := range rows {
		numbers := getLineValues(row)
		list1 = append(list1, numbers[0])
		list2 = append(list2, numbers[1])
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var total float64
	for i := 0; i < len(list1); i++ {
		total += math.Abs(float64(list1[i] - list2[i]))
	}

	return int(total)
}
