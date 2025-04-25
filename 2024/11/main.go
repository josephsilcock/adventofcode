package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	for _, val := range strings.Fields(utils.ReadFile("input.txt")[0]) {
		numbers = append(numbers, utils.ConvertStringToInt(val))
	}

	fmt.Println("Star 1:", getTotalNumbersAfterNIterations(numbers, 25))
	fmt.Println("Star 2:", getTotalNumbersAfterNIterations(numbers, 75))
}

func getTotalNumbersAfterNIterations(numbers []int, iterations int) (total int) {
	cache := make(map[CacheKey]int)

	for _, number := range numbers {
		total += getNumberAfterNIterations(number, iterations, cache)
	}

	return
}

func getNumberAfterNIterations(number int, iterations int, cache map[CacheKey]int) int {
	if iterations == 0 {
		return 1
	}
	cacheKey := CacheKey{number, iterations}
	cachedNumberOfStones, ok := cache[cacheKey]
	if ok {
		return cachedNumberOfStones
	}

	numberOfStones := 0

	if number == 0 {
		numberOfStones = getNumberAfterNIterations(1, iterations-1, cache)
	} else if length := math.Floor(math.Log10(float64(number))) + 1; math.Mod(length, 2) == 0 {
		strVal := strconv.Itoa(number)
		mid := int(length / 2)
		numberOfStones = getNumberAfterNIterations(utils.ConvertStringToInt(strVal[:mid]), iterations-1, cache) + getNumberAfterNIterations(utils.ConvertStringToInt(strVal[mid:]), iterations-1, cache)
	} else {
		numberOfStones = getNumberAfterNIterations(2024*number, iterations-1, cache)
	}

	cache[cacheKey] = numberOfStones
	return numberOfStones
}
