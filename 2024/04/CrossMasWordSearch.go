package main

type CrossMasWordSearch struct {
	WordSearch
}

func (cmws *CrossMasWordSearch) FindCountOfCrossMas() (total int) {
	for rowIndex, letters := range cmws.allLetters {
		for colIndex, letter := range letters {
			if letter == 'A' {
				positionOfA := Point{colIndex, rowIndex}
				if cmws.checkIfAHasCrossMas(positionOfA, Point{1, 1}) && cmws.checkIfAHasCrossMas(positionOfA, Point{-1, 1}) {
					total++
				}
			}
		}
	}
	return
}

func (cmws *CrossMasWordSearch) checkIfAHasCrossMas(positionOfA, direction Point) bool {
	if cmws.getLetter(positionOfA.Add(direction)) == 'M' && cmws.getLetter(positionOfA.Add(direction.Scale(-1))) == 'S' {
		return true
	}
	return cmws.getLetter(positionOfA.Add(direction)) == 'S' && cmws.getLetter(positionOfA.Add(direction.Scale(-1))) == 'M'
}
