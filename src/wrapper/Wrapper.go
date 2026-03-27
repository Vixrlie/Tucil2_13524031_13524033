package wrapper

import (
	"math"
	"tucil/src/octree"
	"tucil/src/point"
)

func WrapInBox(arrPoint []point.Point) (oct *octree.OctreeNode) {
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

	var sd, cx, cy, cz float32 //s: panjang, c: tengah, sd: stengah panjang maks
	// ceil biar ga gmpg error
	cx = float32(math.Ceil(float64((max_x + min_x) / 2)))
	cy = float32(math.Ceil(float64((max_y + min_y) / 2)))
	cz = float32(math.Ceil(float64((max_z + min_z) / 2)))

	var maxDist float32 = 0
	for i := range arrPoint {
		dx := float32(math.Abs(float64(arrPoint[i].X - cx)))
		dy := float32(math.Abs(float64(arrPoint[i].Y - cy)))
		dz := float32(math.Abs(float64(arrPoint[i].Z - cz)))

		if dx > maxDist {
			maxDist = dx
		}
		if dy > maxDist {
			maxDist = dy
		}
		if dz > maxDist {
			maxDist = dz
		}
	}
	sd = maxDist

	// biar kelipatan dua aj
	var exp float64 = 0
	for float32(math.Pow(2, exp)) < sd {
		exp++
	}
	sd = float32(math.Pow(2, exp))

	return octree.NewOctreeNode(point.Point{X: cx, Y: cy, Z: cz}, sd)
}
