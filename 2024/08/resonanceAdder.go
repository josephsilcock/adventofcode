package main

type ResonanceAdder struct{}

func (ra ResonanceAdder) AddAntenna(radarType rune, point Point, am *AntennaeMap) {
	for _, existingPoint := range am.antennaeMapping[radarType] {
		difference := existingPoint.Difference(point)
		for _, direction := range [2]int{-1, 1} {
			antinode := existingPoint
			for am.pointInBounds(antinode) {
				am.antinodeMapping[antinode] = true
				antinode = antinode.Add(difference.Scale(direction))
			}
		}
	}
	am.antennaeMapping[radarType] = append(am.antennaeMapping[radarType], point)
}
