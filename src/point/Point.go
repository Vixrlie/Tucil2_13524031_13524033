package point

import (
	"fmt"
	"os"
)

type Point struct {
	X float32
	Y float32
	Z float32
}

func ToPoint(vert [][3]float32) []Point {
	arrP := make([]Point, len(vert))
	for v := range vert {
		arrP = append(arrP, Point{vert[v][0], vert[v][1], vert[v][2]})

	}
	return arrP
}

func PrintPoints(arrP []Point) {
	for v := range arrP {
		fmt.Printf("%d : %f %f %f\n", v, arrP[v].X, arrP[v].Y, arrP[v].Z)
	}
}

func FPrintPoints(F *os.File, arrP []Point) {
	for v := range arrP {
		fmt.Fprintf(F, "v %f %f %f\n", arrP[v].X, arrP[v].Y, arrP[v].Z)
	}
}
