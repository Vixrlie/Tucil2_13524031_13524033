package voxelizer

import (
	"tucil/src/point"
)

func abs(x float32) float32 {
	if x < 0 { return -x }
	return x
}

func vecSub(a point.Point, b point.Point) point.Point {
	return point.Point{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func vecCross(a point.Point, b point.Point) point.Point {
	return point.Point{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func vecDot(a point.Point, b point.Point) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func AabbTest(centre point.Point, halfSide float32, f point.Face) bool {
	p0 := vecSub(f.V0, centre)
	p1 := vecSub(f.V1, centre)
	p2 := vecSub(f.V2, centre)

	minF_x := min(p0.X, p1.X, p2.X)
	minF_y := min(p0.Y, p1.Y, p2.Y)
	minF_z := min(p0.Z, p1.Z, p2.Z)
	maxF_x := max(p0.X, p1.X, p2.X)
	maxF_y := max(p0.Y, p1.Y, p2.Y)
	maxF_z := max(p0.Z, p1.Z, p2.Z)

	// == Test xyz axes (cek klo si segitiga tbtb jadi kotak gitu)
	if (minF_x > halfSide || maxF_x < -halfSide) { return false }
	if (minF_y > halfSide || maxF_y < -halfSide) { return false }
	if (minF_z > halfSide || maxF_z < -halfSide) { return false }

	// == Test face's normal (cek si kotak nempel/nembus ga ama segitiga)
	l0 := vecSub(p1, p0)
	l1 := vecSub(p2, p1)
	l2 := vecSub(p0, p2)

	n := vecCross(l0, l1)
	cf := vecDot(n, p0) // centre to face
	cc := halfSide * (abs(n.X) + abs(n.Y) + abs(n.Z)) // centre to corner

	if (cf > cc || cf < -cc) { return false }

	// == Test sisanya (smua cross biar gd gap yg miss)
	var min_p, max_p, r float32

	// line 0
	min_p = min(p0.Y*l0.Z - p0.Z*l0.Y, p2.Y*l0.Z - p2.Z*l0.Y) // x & l0
	max_p = max(p0.Y*l0.Z - p0.Z*l0.Y, p2.Y*l0.Z - p2.Z*l0.Y)
	r = halfSide * (abs(l0.Z) + abs(l0.Y))
	if min_p > r || max_p < -r { return false }

	min_p = min(p0.Z*l0.X - p0.X*l0.Z, p2.Z*l0.X - p2.X*l0.Z) // y & l0
	max_p = max(p0.Z*l0.X - p0.X*l0.Z, p2.Z*l0.X - p2.X*l0.Z)
	r = halfSide * (abs(l0.Z) + abs(l0.X))
	if min_p > r || max_p < -r { return false }

	min_p = min(p1.X*l0.Y - p1.Y*l0.X, p2.X*l0.Y - p2.Y*l0.X) // z & l0
	max_p = max(p1.X*l0.Y - p1.Y*l0.X, p2.X*l0.Y - p2.Y*l0.X)
	r = halfSide * (abs(l0.Y) + abs(l0.X))
	if min_p > r || max_p < -r { return false }

	// line 1
	min_p = min(p0.Y*l1.Z - p0.Z*l1.Y, p1.Y*l1.Z - p1.Z*l1.Y) // x & l1
	max_p = max(p0.Y*l1.Z - p0.Z*l1.Y, p1.Y*l1.Z - p1.Z*l1.Y)
	r = halfSide * (abs(l1.Z) + abs(l1.Y))
	if min_p > r || max_p < -r { return false }

	min_p = min(p0.Z*l1.X - p0.X*l1.Z, p1.Z*l1.X - p1.X*l1.Z) // y & l1
	max_p = max(p0.Z*l1.X - p0.X*l1.Z, p1.Z*l1.X - p1.X*l1.Z)
	r = halfSide * (abs(l1.Z) + abs(l1.X))
	if min_p > r || max_p < -r { return false }

	min_p = min(p0.X*l1.Y - p0.Y*l1.X, p1.X*l1.Y - p1.Y*l1.X) // z & l1
	max_p = max(p0.X*l1.Y - p0.Y*l1.X, p1.X*l1.Y - p1.Y*l1.X)
	r = halfSide * (abs(l1.Y) + abs(l1.X))
	if min_p > r || max_p < -r { return false }

	// line 2
	min_p = min(p0.Y*l2.Z - p0.Z*l2.Y, p1.Y*l2.Z - p1.Z*l2.Y) // x & l2
	max_p = max(p0.Y*l2.Z - p0.Z*l2.Y, p1.Y*l2.Z - p1.Z*l2.Y)
	r = halfSide * (abs(l2.Z) + abs(l2.Y))
	if min_p > r || max_p < -r { return false }

	min_p = min(p0.Z*l2.X - p0.X*l2.Z, p1.Z*l2.X - p1.X*l2.Z) // y & l2
	max_p = max(p0.Z*l2.X - p0.X*l2.Z, p1.Z*l2.X - p1.X*l2.Z)
	r = halfSide * (abs(l2.Z) + abs(l2.X))
	if min_p > r || max_p < -r { return false }

	min_p = min(p0.X*l2.Y - p0.Y*l2.X, p1.X*l2.Y - p1.Y*l2.X) // z & l2
	max_p = max(p0.X*l2.Y - p0.Y*l2.X, p1.X*l2.Y - p1.Y*l2.X)
	r = halfSide * (abs(l2.Y) + abs(l2.X))
	if min_p > r || max_p < -r { return false }

	return true
}