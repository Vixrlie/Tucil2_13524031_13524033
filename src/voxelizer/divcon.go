package voxelizer

import (
	"fmt"
	"math"
	"tucil/src/octree"
	"tucil/src/point"
	"tucil/src/wrapper"
)

var DepthList = make(map[int]int)

func StartVoxelize(points []point.Point, faces []point.Face) *octree.OctreeNode {
	_, root := wrapper.WrapInBox(points)

	exp := math.Log2(float64(root.HalfSide)*2)-8.0
	targetSize := math.Pow(2, exp)
	fmt.Println("\nEvery voxel is represented with the size of 2 to the power of 'i'")
	fmt.Println("Please enter your desired 'i'")
	fmt.Println("Hint : 'i' can be negative")
	fmt.Printf("Hint : default 'i' is %v\n", int(exp))
	fmt.Print(">> ")
	fmt.Scanf("%f\n", &exp)

	root.InFaces = faces
	Divide(root, 0, targetSize)
	return root
}

func Divide(node *octree.OctreeNode, depth int, targetSize float64) {
	DepthList[depth]++

	if node.HalfSide <= float32(targetSize)/2 {
		node.IsLeaf = true
		return
	}

	// make children
	childHalf := node.HalfSide / 2.0
	offsets := [8][3]float32{
		{1, 1, 1}, {-1, 1, 1}, {1, -1, 1}, {-1, -1, 1},
		{1, 1, -1}, {-1, 1, -1}, {1, -1, -1}, {-1, -1, -1},
	}

	// iterate all children
	for i := 0; i < 8; i++ {
		childCentre := point.Point{
			X: node.Centre.X + (offsets[i][0] * childHalf),
			Y: node.Centre.Y + (offsets[i][1] * childHalf),
			Z: node.Centre.Z + (offsets[i][2] * childHalf),
		}

		var childInclFaces []point.Face

		// test for any intersection
		for _, f := range node.InFaces {
			if AabbTest(childCentre, childHalf, f) {
				childInclFaces = append(childInclFaces, f)
			}
		}

		// make real children
		if len(childInclFaces) > 0 {
			node.Children[i] = &octree.OctreeNode{
				Centre:   childCentre,
				HalfSide: childHalf,
				InFaces:  childInclFaces,
				IsLeaf:   false,
			}
		}
	}

	node.InFaces = nil

	for i := 0; i < 8; i++ {
		if node.Children[i] != nil { Divide(node.Children[i], depth+1, targetSize) }
	}
}