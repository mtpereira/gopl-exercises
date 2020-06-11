// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	run()
}

func run() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, dc := corner(i+1, j+1)
			s := ax + ay + bx + by + cx + cy + dx + dy
			if math.IsNaN(s) || math.IsInf(s, 1) || math.IsInf(s, -1) {
				continue
			}
			fmt.Printf("<polygon stroke='%v' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				dc, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy float64, colour string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := eggBox(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	if z >= 0 {
		colour = "#ff0000"
		return
	}
	colour = "#0000ff"
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	a, b := 0.25, 0.7
	return a * (math.Sin(x/b) + math.Sin(y/b))
}

func saddle(x, y float64) float64 {
	return math.Pow(x, 2) - math.Pow(y, 2)
}

//!-
