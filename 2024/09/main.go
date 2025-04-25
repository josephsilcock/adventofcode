package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
	"math"
)

func main() {
	fileMap := []rune(utils.ReadFile("input.txt")[0])
	if math.Mod(float64(len(fileMap)), 2) == 0 {
		fileMap = fileMap[:len(fileMap)-1]
	}

	fileMapSizes := make([]int, len(fileMap))

	for i, file := range fileMap {
		fileMapSizes[i] = utils.ConvertStringToInt(string(file))
	}

	star1 := getChecksumWithFragmentation(fileMapSizes)
	star2 := getChecksumWithNoFragmentation(fileMapSizes)

	fmt.Println("Star 1:", star1)
	fmt.Println("Star 2:", star2)
}
