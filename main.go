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

func main() {
	p := Point{x: 1, y: 2}
	fmt.Println(p)
	p.move(4, 5)
	fmt.Println(p)

}
