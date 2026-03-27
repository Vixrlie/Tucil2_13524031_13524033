package octree

import (
	"tucil/src/point"
)

type OctreeNode struct {
	Children [8]*OctreeNode
	InFaces  []point.Face

	Centre   point.Point // Titik tengah kubus
	HalfSide float32     // Setengah panjang sisi kubus

	IsLeaf bool
}

func NewOctreeNode(p point.Point, s float32) *OctreeNode {
	return &OctreeNode{
		Centre:   point.Point{X: p.X, Y: p.Y, Z: p.Z},
		HalfSide: s,
		IsLeaf:   false,
	}
}
