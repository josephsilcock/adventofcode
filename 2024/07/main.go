package main

import (
	"fmt"
	"github.com/josephsilcock/adventofcode/2024/utils"
	"sync"
)

func main() {
	rows := utils.ReadFile("input.txt")
	totalCalibration, totalCalibrationIncludingCombination := processRows(rows)

	fmt.Println("Star 1:", totalCalibration)
	fmt.Println("Star 1:", totalCalibrationIncludingCombination)
}

func processRows(rows []string) (totalCalibration int, totalCalibrationIncludingCombination int) {
	var wg sync.WaitGroup
	calibrationCh := make(chan int)
	calibrationIncludingCombinationCh := make(chan int)

	for _, row := range rows {
		target, values := parseInputLine(row)
		wg.Add(2)

		go func() {
			if isCalculationPossible(target, values) {
				calibrationCh <- target
			}
			wg.Done()
		}()
		go func() {
			if isCalculationPossibleIncludingCombination(target, values) {
				calibrationIncludingCombinationCh <- target
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(calibrationCh)
		close(calibrationIncludingCombinationCh)
	}()

	done := make(chan struct{})
	go func() {
		for calibration := range calibrationCh {
			totalCalibration += calibration
		}
		done <- struct{}{}
	}()

	for calibration := range calibrationIncludingCombinationCh {
		totalCalibrationIncludingCombination += calibration
	}
	<-done

	return
}
