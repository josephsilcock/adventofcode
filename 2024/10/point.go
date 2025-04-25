package main

type Point struct{ x, y int }

func (point Point) Add(otherPoint Point) Point {
	point.x += otherPoint.x
	point.y += otherPoint.y
	return point
}
