package wrapper

import (
	"math"
	"tucil/src/octree"
	"tucil/src/point"
)

func WrapInBox(arrPoint []point.Point) (res []point.Point, oct *octree.OctreeNode) {
	oct = nil
	var max_x, max_y, max_z float32 = -math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32

	var min_x, min_y, min_z float32 = -max_x, -max_y, -max_z

	// cari titik titik minimum maksimum untuk dibuat balok yang mengelilingi
	for v := range arrPoint {
		x := arrPoint[v].X
		y := arrPoint[v].Y
		z := arrPoint[v].Z

		if x > max_x {
			max_x = x
		}
		if x < min_x {
			min_x = x
		}

		if y > max_y {
			max_y = y
		}
		if y < min_y {
			min_y = y
		}

		if z > max_z {
			max_z = z
		}
		if z < min_z {
			min_z = z
		}
	}

	var sx, sy, sz, sd, cx, cy, cz float32 //s: panjang, c: tengah, sd: panjang maks
	sx = max_x - min_x
	cx = (max_x + min_x) / 2
	sy = max_y - min_y
	cy = (max_y + min_y) / 2
	sz = max_z - min_z
	cz = (max_z + min_z) / 2

	sd = sx
	if sy > sd {
		sd = sy
	}
	if sz > sd {
		sd = sz
	}

	min_x = cx - sd/2
	max_x = cx + sd/2
	min_y = cy - sd/2
	max_y = cy + sd/2
	min_z = cz - sd/2
	max_z = cz + sd/2

	res = []point.Point{}

	res = append(res, point.Point{min_x, min_y, min_z},
		point.Point{min_x, min_y, max_z},
		point.Point{min_x, max_y, min_z},
		point.Point{min_x, max_y, max_z},
		point.Point{max_x, min_y, min_z},
		point.Point{max_x, min_y, max_z},
		point.Point{max_x, max_y, min_z},
		point.Point{max_x, max_y, max_z})

	return res, octree.NewOctreeNode(point.Point{cx, cy, cz}, sd)
}
