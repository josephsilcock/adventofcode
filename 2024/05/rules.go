package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getRules() map[string][]string {
	rules := make(map[string][]string)

	file, err := os.Open("input_rules.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "|")

		rules[numbers[0]] = append(rules[numbers[0]], numbers[1])
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules
}
