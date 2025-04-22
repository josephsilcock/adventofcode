package main

import (
	"log"
	"strconv"
	"strings"
)

func getLineValues(line string) (ret [2]int) {
	parts := strings.Fields(line)

	firstValue, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	ret[0] = firstValue

	secondValue, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	ret[1] = secondValue

	return
}
