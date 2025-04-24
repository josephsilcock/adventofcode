package main

type Point struct{ x, y int }

func (point Point) Add(otherPoint Point) Point {
	point.x += otherPoint.x
	point.y += otherPoint.y
	return point
}

func (point Point) Scale(scaler int) Point {
	point.x *= scaler
	point.y *= scaler
	return point
}

func (point Point) Difference(otherPoint Point) Point {
	return Point{x: otherPoint.x - point.x, y: otherPoint.y - point.y}
}
