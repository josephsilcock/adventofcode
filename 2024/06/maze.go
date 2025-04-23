package main

type FacingDirection Point

var (
	FacingUp    = FacingDirection{0, -1}
	FacingRight = FacingDirection{1, 0}
	FacingDown  = FacingDirection{0, 1}
	FacingLeft  = FacingDirection{-1, 0}
)

type Maze struct {
	startingLocation Point
	currentLocation  Point
	direction        FacingDirection
	positions        [][]Position
	visitedPositions int
}

func (m *Maze) AddRow(row string) {
	var rowPositions []Position
	for _, character := range []rune(row) {
		switch character {
		case '.':
			rowPositions = append(rowPositions, Position{})
		case '#':
			rowPositions = append(rowPositions, Position{obstacle: true})
		case '^':
			m.startingLocation = Point{len(rowPositions), len(m.positions)}
			rowPositions = append(rowPositions, Position{})
		}
	}
	m.positions = append(m.positions, rowPositions)
}

func (m *Maze) visitPosition(point Point) {
	position := m.GetPositionAtPoint(point)
	if len(position.visitedDirections) == 0 {
		m.visitedPositions++
	}
	m.currentLocation = point
	position.visitedDirections = append(position.visitedDirections, m.direction)
}

func (m *Maze) rotate() {
	switch m.direction {
	case FacingUp:
		m.direction = FacingRight
	case FacingRight:
		m.direction = FacingDown
	case FacingDown:
		m.direction = FacingLeft
	case FacingLeft:
		m.direction = FacingUp
	}
	currentPosition := m.GetPositionAtPoint(m.currentLocation)
	currentPosition.visitedDirections = append(currentPosition.visitedDirections, m.direction)

}

func (m *Maze) WillLeaveBounds() bool {
	newPoint := m.currentLocation.Add(Point(m.direction))
	return newPoint.x < 0 || newPoint.x >= len(m.positions[0]) || newPoint.y < 0 || newPoint.y >= len(m.positions)
}

func (m *Maze) WillHaveBeenHereBefore() bool {
	newPoint := m.currentLocation.Add(Point(m.direction))
	newPosition := m.GetPositionAtPoint(newPoint)
	for _, direction := range newPosition.visitedDirections {
		if direction == m.direction {
			return true
		}
	}
	return false
}

func (m *Maze) Move() {
	newPoint := m.currentLocation.Add(Point(m.direction))
	newPosition := m.GetPositionAtPoint(newPoint)
	if newPosition.obstacle {
		m.rotate()
	} else {
		m.visitPosition(newPoint)
	}
}

func (m *Maze) StartMaze() {
	m.direction = FacingUp
	m.visitPosition(m.startingLocation)
}

func (m *Maze) GetPositionAtPoint(point Point) *Position {
	return &m.positions[point.y][point.x]
}

func (m *Maze) CreateCopy() *Maze {
	positionsCopy := make([][]Position, len(m.positions))
	for i := range m.positions {
		rowCopy := make([]Position, len(m.positions[i]))
		copy(rowCopy, m.positions[i])
		positionsCopy[i] = rowCopy
	}

	return &Maze{startingLocation: m.startingLocation, positions: positionsCopy}
}
