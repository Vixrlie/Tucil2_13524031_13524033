package wrapper

import (
	"math"
	"tucil/src/point"
)

func WrapInBox(arrPoint []point.Point) (res []point.Point) {
	var max_x, max_y, max_z float32 = -math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32

	var min_x, min_y, min_z float32 = -max_x, -max_y, -max_z

	// cari titik titik minimum maksimum untuk dibuat balok yang mengelilingi
	for v := range arrPoint {
		x := arrPoint[v].X
		y := arrPoint[v].Y
		z := arrPoint[v].Z

		if x > max_x {
			max_x = x
		} else if x < min_x {
			min_x = x
		}

		if y > max_y {
			max_y = y
		} else if y < min_y {
			min_y = y
		}

		if z > max_z {
			max_z = z
		} else if z < min_z {
			min_z = z
		}
	}

	res = []point.Point{}

	res = append(res, point.Point{min_x, min_y, min_z},
		point.Point{min_x, min_y, max_z},
		point.Point{min_x, max_y, min_z},
		point.Point{min_x, max_y, max_z},
		point.Point{max_x, min_y, min_z},
		point.Point{max_x, min_y, max_z},
		point.Point{max_x, max_y, min_z},
		point.Point{max_x, max_y, max_z})

	return res
}
