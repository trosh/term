package main

import (
	"fmt"
	"time"
	"math"
	"github.com/trosh/term-go/term"
)

func Triangle(center pxl,
              radius float64,
              angle  float64) (pxl,
                               pxl,
                               pxl) {
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
		int(math.Cos(angle3)*radius) + center.y}
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
