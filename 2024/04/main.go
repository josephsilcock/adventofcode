package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	xmasWordSearch := XmasWordSearch{}
	crossMasWordSearch := CrossMasWordSearch{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xmasWordSearch.AddRow(scanner.Text())
		crossMasWordSearch.AddRow(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Star 1:", xmasWordSearch.FindCountOfXmas())
	fmt.Println("Star 2:", crossMasWordSearch.FindCountOfCrossMas())
}
