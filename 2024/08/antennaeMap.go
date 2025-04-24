package main

type AntennaAdder interface {
	AddAntenna(radarType rune, point Point, am *AntennaeMap)
}

type AntennaeMap struct {
	height, width   int
	antennaeMapping map[rune][]Point
	antinodeMapping map[Point]bool
	antennaAdder    AntennaAdder
}

func (am *AntennaeMap) pointInBounds(point Point) bool {
	return point.x >= 0 && point.x < am.width && point.y >= 0 && point.y < am.height
}

func (am *AntennaeMap) AddRow(row string, rowIndex int) {
	for columnIndex, character := range []rune(row) {
		if character != '.' {
			newPoint := Point{columnIndex, rowIndex}
			am.antennaAdder.AddAntenna(character, newPoint, am)
		}
	}
}

func (am *AntennaeMap) GetTotalAntinodes() int {
	return len(am.antinodeMapping)
}
