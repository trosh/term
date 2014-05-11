package main

import (
	"fmt"
	"time"
	"math"
)

type pxl struct {
	x int
	y int
}

type scr struct {
	p1 pxl
	p2 pxl
	bg int
}

func (s scr) Size() pxl {
	size := pxl{s.p2.x - s.p1.x, s.p2.y - s.p1.x}
	if s.p2.x < s.p1.x { size.x = size.x }
	if s.p2.y < s.p1.y { size.y = size.y }
	return size
}

func (s scr) Flush() {
	fmt.Printf("\033[48;5;%dm", s.bg)
	for y := s.p1.y; y <= s.p2.y; y += 1 {
		fmt.Printf("\033[%d;%dH", y, s.p1.x)
		for x := s.p1.x; x <= s.p2.x; x += 1 {
			fmt.Printf(" ")
		}
	}
}

func (s scr) Plot(p pxl, bg int) {
	if p.x >= s.p1.x && p.x <= s.p2.x &&
	p.y >= s.p1.y && p.y <= s.p2.y {
		fmt.Printf("\033[%d;%dH\033[48;5;%dm ", p.y, p.x, bg)
	}
}

func (s scr) Line(p1, p2 pxl, color bool) {
	if p1.x > p2.x { //make sure p1 has lowest x (is at the left)
		p1, p2 = p2, p1
	}
	incx := true
	if ( p2.y > p1.y && p2.y - p1.y > p2.x - p1.x ) ||
	   ( p2.y < p1.y && p1.y - p2.y > p2.x - p1.x ) {incx = false}
	if incx {
		slope := float64(p2.y - p1.y) / float64(p2.x - p1.x)
		xstart := s.p1.x
		if p1.x > s.p1.x { xstart = p1.x }
		xend := s.p2.x
		if p2.x < s.p2.x { xend = p2.x }
		for x := xstart; x <= xend; x += 1 {
			y := p1.y + int(slope * float64(x - p1.x) + 0.5)
			p := pxl{x, y}
			if color {
				s.Plot(p, y*6 + 16)
			} else {
				s.Plot(p, 0)
			}
		}
	} else {
		if p1.y > p2.y { //make sure p1 has lowest y (is at the top)
			p1, p2 = p2, p1
		}
		slope := float64(p2.x - p1.x) / float64(p2.y - p1.y)
		ystart := s.p1.y
		if p1.y > s.p1.y { ystart = p1.y }
		yend := s.p2.y
		if p2.y < s.p2.y { yend = p2.y }
		for y := ystart; y <= yend; y += 1 {
			x := p1.x + int(slope * float64(y - p1.y) + 0.5)
			p := pxl{x, y}
			if color {
				s.Plot(p, y*6 + 16)
			} else {
				s.Plot(p, 0)
			}
		}
	}
}

func Triangle(center pxl, radius float64, angle float64) (pxl, pxl, pxl) {
	angle2 := angle + 2 * math.Pi / 3
	angle3 := angle + 4 * math.Pi / 3
	p1 := pxl{
		int(math.Sin(angle)*radius * 1.7) + center.x,
		int(math.Cos(angle)*radius      ) + center.y}
	p2 := pxl{
		int(math.Sin(angle2)*radius * 1.7) + center.x,
		int(math.Cos(angle2)*radius      ) + center.y}
	p3 := pxl{
		int(math.Sin(angle3)*radius * 1.7) + center.x,
		int(math.Cos(angle3)*radius      ) + center.y}
	return p1, p2, p3
}

func main() {
	s := scr{pxl{1, 1}, pxl{60, 28}, 0}
	s.Flush()
	s = scr{pxl{10, 7}, pxl{40, 22}, 52}
	orig := pxl{25, 15}
	radius := 13.0
	for i := 0.0; ; i += 0.1 {
		p1, p2, p3 := Triangle(orig, radius, math.Sin(i) + i)
		s.Line(p1, p2, true)
		s.Line(p2, p3, true)
		s.Line(p3, p1, true)
		fmt.Printf("\033[0;0H")
		time.Sleep(25 * time.Millisecond)
		s.Flush()
	}
	fmt.Printf("\033[0m\n")
}
