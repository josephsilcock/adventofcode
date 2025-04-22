package main

import (
	"bufio"
	"log"
	"os"
)

func getListSimilarity() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	list1 := make(map[int]int)
	list2 := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := getLineValues(scanner.Text())
		list1[numbers[0]]++
		list2[numbers[1]]++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var similarityScore int

	for number, occurrences := range list1 {
		similarityScore += number * occurrences * list2[number]
	}

	return similarityScore
}
