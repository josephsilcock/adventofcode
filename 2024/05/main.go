package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
)

func main() {
	rules := getRules()
	orderedUpdates, unorderedUpdates := categoriseUpdates(rules)

	fmt.Println("Star 1:", sumMiddleValues(orderedUpdates))
	fmt.Println("Star 2:", sumMiddleValues(orderUpdates(unorderedUpdates, rules)))
}

func sumMiddleValues(updates [][]string) (total int) {
	for _, update := range updates {
		middlePageIndex := (len(update) - 1) / 2
		middleValue := utils.ConvertStringToInt(update[middlePageIndex])

		total += middleValue
	}
	return
}

func orderUpdates(updates [][]string, rules map[string][]string) (orderedUpdates [][]string) {
	for _, update := range updates {
		reorderedUpdate := update
		for {
			isOrdered, firstIndex, secondIndex := isUpdateOrdered(reorderedUpdate, rules)
			if isOrdered {
				break
			}
			reorderedUpdate[firstIndex], reorderedUpdate[secondIndex] = reorderedUpdate[secondIndex], reorderedUpdate[firstIndex]
		}
		orderedUpdates = append(orderedUpdates, reorderedUpdate)
	}

	return
}
