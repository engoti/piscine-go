package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Printnum(num int) {
	a := '0'
	for i := 0; i < num; i++ {
		a++
	}
	z01.PrintRune(a)
}

func Printstr(s string) {
	for _, arg := range s {
		z01.PrintRune(arg)
	}
}

func main() {
	part1 := "x = "
	part2 := ", y = "
	points := &point{}
	setPoint(points)
	Printstr(part1)
	Printnum(points.x / 10)
	Printnum(points.x % 10)
	Printstr(part2)
	Printnum(points.y / 10)
	Printnum(points.y % 10)

	z01.PrintRune('\n')
}
