package point

import (
	"fmt"
)

type Face struct {
	V0 Point
	V1 Point
	V2 Point
}

func ToFace(p []Point, list [][3]int) []Face {
	res := make([]Face, 0, len(list))
	for _, f := range list {
		temp := Face{ p[f[0]-1],
					  p[f[1]-1], 
					  p[f[2]-1], }
		res = append(res, temp)
	}
	return res
}

func PrintFaces(list []Face) {
	for i, f := range list {
		fmt.Printf("Face - %v\n", i+1)
		fmt.Printf("V0: %+v\n", f.V0)
		fmt.Printf("V1: %+v\n", f.V1)
		fmt.Printf("V1: %+v\n", f.V2)
		fmt.Println()
	}
}