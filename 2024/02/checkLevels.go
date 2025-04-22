package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkLevels(allowOneBadLevel bool) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var safeReports int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		level := Level{canSkipValue: allowOneBadLevel}
		for _, valueText := range strings.Fields(scanner.Text()) {
			value, err := strconv.Atoi(valueText)
			if err != nil {
				log.Fatal(err)
			}
			level.AddValue(value)
		}

		if level.IsSafe() {
			safeReports++
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return safeReports
}
