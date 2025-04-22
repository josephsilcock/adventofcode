package main

import (
	"math"
)

type LevelType int

const (
	LevelUnknow LevelType = iota
	LevelAscending
	LevelDescending
)

type Level struct {
	levelType    LevelType
	canSkipValue bool
	values       []int
	unsafe       bool
}

func (level *Level) checkValueIsSafe(value int) bool {
	if len(level.values) == 0 {
		return true
	}

	lastValue := level.values[len(level.values)-1]
	difference := math.Abs(float64(lastValue - value))

	if difference == 0 || difference > 3 {
		return false
	}

	var newLevelType LevelType

	if value > lastValue {
		newLevelType = LevelAscending
	} else {
		newLevelType = LevelDescending
	}

	sameLevelType := level.levelType == LevelUnknow || newLevelType == level.levelType

	level.levelType = newLevelType

	return sameLevelType
}

func (level *Level) AddValue(value int) {
	if !level.checkValueIsSafe(value) {
		level.unsafe = true
	}
	level.values = append(level.values, value)
}

func (level *Level) createCopyMissingIndex(indexToSkip int) *Level {
	newLevel := &Level{canSkipValue: false}

	for index, value := range level.values {
		if index != indexToSkip {
			newLevel.AddValue(value)
		}
	}

	return newLevel
}

func (level *Level) IsSafe() bool {
	if !level.unsafe {
		return true
	}
	if !level.canSkipValue {
		return false
	}
	for index := range level.values {
		if level.createCopyMissingIndex(index).IsSafe() {
			return true
		}
	}
	return false
}
