package main

var (
	xmasDirections = [8]Point{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
)

type XmasWordSearch struct {
	WordSearch
}

func (cmws *XmasWordSearch) FindCountOfXmas() (total int) {
	for rowIndex, letters := range cmws.allLetters {
		for colIndex, letter := range letters {
			if letter == 'X' {
				positionOfX := Point{colIndex, rowIndex}
				for _, direction := range xmasDirections {
					if cmws.checkIfXHasXmas(positionOfX, direction) {
						total++
					}
				}
			}
		}
	}
	return
}

func (cmws *XmasWordSearch) checkIfXHasXmas(positionOfX, direction Point) bool {
	if cmws.getLetter(positionOfX.Add(direction)) != 'M' {
		return false
	}
	if cmws.getLetter(positionOfX.Add(direction.Scale(2))) != 'A' {
		return false
	}
	return cmws.getLetter(positionOfX.Add(direction.Scale(3))) == 'S'
}
