package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
)

func getListDistance() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := getLineValues(scanner.Text())
		list1 = append(list1, numbers[0])
		list2 = append(list2, numbers[1])
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var total float64
	for i := 0; i < len(list1); i++ {
		total += math.Abs(float64(list1[i] - list2[i]))
	}

	return int(total)
}
