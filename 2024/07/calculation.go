package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"math"
	"strconv"
)

func isCalculationPossible(target int, values []int) bool {
	operatorRepresentation := 0
	operatorRepresentationCap := int(math.Pow(2, float64(len(values)-1)))

OPERATORS:
	for operatorRepresentation < operatorRepresentationCap {
		currentTotal := values[0]
		currentBit := 0
		for _, value := range values[1:] {
			if operatorRepresentation&(1<<currentBit) == 0 {
				currentTotal *= value
			} else {
				currentTotal += value
			}
			if currentTotal > target {
				operatorRepresentation++
				continue OPERATORS
			}
			currentBit++
		}
		if currentTotal == target {
			return true
		}
		operatorRepresentation++
	}

	return false
}

func isCalculationPossibleIncludingCombination(target int, values []int) bool {
	operatorRepresentation := 0
	operatorRepresentationCap := int(math.Pow(3, float64(len(values)-1)))

OPERATORS:
	for operatorRepresentation < operatorRepresentationCap {
		currentTotal := values[0]
		currentBit := 0
		ternary := intToTernary(operatorRepresentation)
		for _, value := range values[1:] {
			bitValue := 0
			if len(ternary) > currentBit {
				bitValue = ternary[currentBit]
			}
			switch bitValue {
			case 0:
				currentTotal *= value
			case 1:
				currentTotal += value
			case 2:
				currentTotal = utils.ConvertStringToInt(strconv.Itoa(currentTotal) + strconv.Itoa(value))
			}
			if currentTotal > target {
				operatorRepresentation++
				continue OPERATORS
			}
			currentBit++
		}
		if currentTotal == target {
			return true
		}
		operatorRepresentation++
	}

	return false
}

func intToTernary(number int) (result []int) {
	for number > 0 {
		remainder := number % 3
		result = append(result, remainder)
		number = number / 3
	}
	return
}
