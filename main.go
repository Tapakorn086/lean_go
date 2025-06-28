package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p *Point) move(x, y int) {
	p.x += x
	p.y += y
}

func (p *Point) display() string{
	return fmt.Sprintf("Point: (%d, %d)", p.x, p.y)
}

type Car struct{
	Point
	Color string
}

func main() {
	p := Point{x: 1, y: 2}
	fmt.Println(p)
	p.move(4, 5)
	fmt.Println(p)

	car := Car{Color:"red",Point: Point{x:5,y:10}}
	car.move(2,3)
	fmt.Println(car.display())

}
