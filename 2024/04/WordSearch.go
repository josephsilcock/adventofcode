package main

type WordSearch struct {
	allLetters [][]rune
}

func (ws *WordSearch) AddRow(row string) {
	ws.allLetters = append(ws.allLetters, []rune(row))
}

func (ws *WordSearch) getLetter(point Point) rune {
	if point.x < 0 || point.x >= len(ws.allLetters[0]) || point.y < 0 || point.y >= len(ws.allLetters) {
		return 0
	}
	return ws.allLetters[point.y][point.x]
}
