package main

import (
	"fmt"
	"time"
	"math"
	"strings"
)

type pxl struct {
	x int
	y int
}

func Plot(p pxl, bg int, size pxl) {
	if p.x >= 1 && p.x <= size.x &&
	p.y >= 1 && p.y <= size.y {
		fmt.Printf("\033[%d;%dH\033[48;5;%dm ", p.y, p.x, bg)
	}
}

func Line(p1, p2 pxl, color bool, size pxl) {
	if p1.x > p2.x { //make sure p1 has lowest x (is at the left)
		p1, p2 = p2, p1
	}
	incx := true
	if ( p2.y > p1.y && p2.y - p1.y > p2.x - p1.x ) ||
	   ( p2.y < p1.y && p1.y - p2.y > p2.x - p1.x ) {incx = false}
	if incx {
		slope := float64(p2.y - p1.y) / float64(p2.x - p1.x)
		xstart := 1
		if p1.x > 1 { xstart = p1.x }
		xend := size.x
		if p2.x < size.x { xend = p2.x }
		for x := xstart; x <= xend; x += 1 {
			y := p1.y + int(slope * float64(x - p1.x) + 0.5)
			p := pxl{x, y}
			if color {
				Plot(p, y*6 + 16, size)
			} else {
				Plot(p, 0, size)
			}
		}
	} else {
		if p1.y > p2.y { //make sure p1 has lowest y (is at the top)
			p1, p2 = p2, p1
		}
		slope := float64(p2.x - p1.x) / float64(p2.y - p1.y)
		ystart := 1
		if p1.y > 1 { ystart = p1.y }
		yend := size.y
		if p2.y < size.y { yend = p2.y }
		for y := ystart; y <= yend; y += 1 {
			x := p1.x + int(slope * float64(y - p1.y) + 0.5)
			p := pxl{x, y}
			if color {
				Plot(p, y*6 + 16, size)
			} else {
				Plot(p, 0, size)
			}
		}
	}
}

func Triangle(center pxl, radius float64, angle float64) (pxl, pxl, pxl) {
	angle2 := angle + 2 * math.Pi / 3
	angle3 := angle + 4 * math.Pi / 3
	p1 := pxl{
		int(math.Sin(angle)*radius * 1.7) + center.x,
		int(math.Cos(angle)*radius) + center.y}
	p2 := pxl{
		int(math.Sin(angle2)*radius * 1.7) + center.x,
		int(math.Cos(angle2)*radius) + center.y}
	p3 := pxl{
		int(math.Sin(angle3)*radius * 1.7) + center.x,
		int(math.Cos(angle3)*radius) + center.y}
	return p1, p2, p3
}

func main() {
	size := pxl{50, 34}
	cleanstr := ""
	for i := 0; i < size.y; i++ {
		cleanstr += strings.Repeat(" ", size.x)
		cleanstr += "\n"
	}
	orig := pxl{25, 15}
	radius := 13.0
	for i := 0.0; ; i += 0.1 {
		p1, p2, p3 := Triangle(orig, radius, math.Sin(i) + i)
		Line(p1, p2, true, size)
		Line(p2, p3, true, size)
		Line(p3, p1, true, size)
		fmt.Printf("\033[0;0H")
		time.Sleep(25 * time.Millisecond)
		fmt.Printf("%s", cleanstr)
	}
	fmt.Printf("\033[0m\n")
}
