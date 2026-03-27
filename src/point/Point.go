package point

import (
	"fmt"
)

type Point struct {
	X float32
	Y float32
	Z float32
}

func ToPoint(vert [][3]float32) []Point {
	arrP := make([]Point, len(vert))
	for i, v := range vert {
		arrP[i] = Point{v[0], v[1], v[2]}
	}
	return arrP
}

func PrintPoints(arrP []Point) {
	for v := range arrP {
		fmt.Printf("%d : %.6f %.6f %.6f\n", v, arrP[v].X, arrP[v].Y, arrP[v].Z)
	}
}
