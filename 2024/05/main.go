package main

import (
	"fmt"
	"log"
	"strconv"
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
		middleValue, err := strconv.Atoi(update[middlePageIndex])
		if err != nil {
			log.Fatal(err)
		}

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
