package octree

import (
	"tucil/src/point"
)

type OctreeNode struct {
	X, Y, Z float32 //Titik tengah kubus
	S       float32 //Panjang sisi kubus

	IsLeaf    bool
	IsSurface bool

	Children [8]*OctreeNode
}

func NewOctreeNode(p point.Point, s float32) *OctreeNode {
	return &OctreeNode{
		X:         p.X,
		Y:         p.Y,
		Z:         p.Z,
		S:         s,
		IsLeaf:    false,
		IsSurface: false,
	}
}
