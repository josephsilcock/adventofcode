package main

type NoResonanceAdder struct{}

func (nra NoResonanceAdder) AddAntenna(radarType rune, point Point, am *AntennaeMap) {
	for _, existingPoint := range am.antennaeMapping[radarType] {
		difference := existingPoint.Difference(point)
		antinodes := [2]Point{point.Add(difference), existingPoint.Add(difference.Scale(-1))}
		for _, antinode := range antinodes {
			if am.pointInBounds(antinode) {
				am.antinodeMapping[antinode] = true
			}
		}
	}
	am.antennaeMapping[radarType] = append(am.antennaeMapping[radarType], point)
}
