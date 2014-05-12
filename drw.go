package term

import (
	"fmt"
	"math"
)

type Pxl struct {
	X int
	Y int
}

type Scr struct {
	P1 Pxl
	P2 Pxl
	Bg int
}

func (s Scr) Size() Pxl {
	size := Pxl{s.P2.X - s.P1.X, s.P2.Y - s.P1.X}
	if s.P2.X < s.P1.X { size.X = size.X }
	if s.P2.Y < s.P1.Y { size.Y = size.Y }
	return size
}

func (s Scr) Flush() {
	fmt.Printf("\033[48;5;%dm", s.Bg)
	for y := s.P1.Y; y <= s.P2.Y; y += 1 {
		fmt.Printf("\033[%d;%dH", y, s.P1.X)
		for x := s.P1.X; x <= s.P2.X; x += 1 {
			fmt.Printf(" ")
		}
	}
}

func (s Scr) Plot(p Pxl, bg int) {
	if p.X >= s.P1.X && p.X <= s.P2.X &&
	p.Y >= s.P1.Y && p.Y <= s.P2.Y {
		fmt.Printf("\033[%d;%dH\033[48;5;%dm ", p.Y, p.X, bg)
	}
}

func (s Scr) Line(p1, p2 Pxl, color bool) {
	if p1.X > p2.X { //make sure p1 has lowest x (is at the left)
		p1, p2 = p2, p1
	}
	incx := true
	if ( p2.Y > p1.Y && p2.Y - p1.Y > p2.X - p1.X ) ||
	   ( p2.Y < p1.Y && p1.Y - p2.Y > p2.X - p1.X ) {incx = false}
	if incx {
		slope := float64(p2.Y - p1.Y) / float64(p2.X - p1.X)
		xstart := s.P1.X
		if p1.X > s.P1.X { xstart = p1.X }
		xend := s.P2.X
		if p2.X < s.P2.X { xend = p2.X }
		for x := xstart; x <= xend; x += 1 {
			y := p1.Y + int(slope * float64(x - p1.X) + 0.5)
			p := Pxl{x, y}
			if color {
				s.Plot(p, y*6 + 16)
			} else {
				s.Plot(p, 0)
			}
		}
	} else {
		if p1.Y > p2.Y { //make sure p1 has lowest y (is at the top)
			p1, p2 = p2, p1
		}
		slope := float64(p2.X - p1.X) / float64(p2.Y - p1.Y)
		ystart := s.P1.Y
		if p1.Y > s.P1.Y { ystart = p1.Y }
		yend := s.P2.Y
		if p2.Y < s.P2.Y { yend = p2.Y }
		for y := ystart; y <= yend; y += 1 {
			x := p1.X + int(slope * float64(y - p1.Y) + 0.5)
			p := Pxl{x, y}
			if color {
				s.Plot(p, y*6 + 16)
			} else {
				s.Plot(p, 0)
			}
		}
	}
}
