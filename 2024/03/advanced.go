package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getAdvancedTotal() (total int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	enabled := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := regex.FindAllStringSubmatch(scanner.Text(), -1)
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

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
