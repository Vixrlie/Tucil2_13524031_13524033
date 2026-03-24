package octree

import (
	"tucil/src/point"
)

type OctreeNode struct {
	Centre  point.Point //Titik tengah kubus
	S       float32 //Panjang sisi kubus

	IsLeaf    bool
	IsSurface bool

	InFaces   []point.Face	  
	Children  [8]*OctreeNode
}

func NewOctreeNode(p point.Point, s float32) *OctreeNode {
	return &OctreeNode{
		Centre:    point.Point{X: p.X, Y: p.Y, Z: p.Z},
		S:         s,
		IsLeaf:    false,
		IsSurface: false,
	}
}
