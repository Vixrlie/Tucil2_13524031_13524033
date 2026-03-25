package voxelizer

import (
	"fmt"
	"os"
	"tucil/src/octree"
)

func GetLeaves(node *octree.OctreeNode, leaves *[]*octree.OctreeNode) {
	if node == nil {
		return
	}

	if node.IsLeaf {
		*leaves = append(*leaves, node)
		return
	}

	for i := range 8 {
		GetLeaves(node.Children[i], leaves)
	}
}

func ExportToOBJ(filename string, leaves []*octree.OctreeNode) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "# Voxels by Vinli Owen\n\n")

	for _, leaf := range leaves {
		cx, cy, cz := leaf.Centre.X, leaf.Centre.Y, leaf.Centre.Z
		h := leaf.HalfSide

		fmt.Fprintf(file, "v %f %f %f\n", cx-h, cy-h, cz+h) // bot l
		fmt.Fprintf(file, "v %f %f %f\n", cx+h, cy-h, cz+h) // bot r
		fmt.Fprintf(file, "v %f %f %f\n", cx+h, cy+h, cz+h) // top r
		fmt.Fprintf(file, "v %f %f %f\n", cx-h, cy+h, cz+h) // top l

		fmt.Fprintf(file, "v %f %f %f\n", cx-h, cy-h, cz-h) // bot l
		fmt.Fprintf(file, "v %f %f %f\n", cx+h, cy-h, cz-h) // bot r
		fmt.Fprintf(file, "v %f %f %f\n", cx+h, cy+h, cz-h) // top r
		fmt.Fprintf(file, "v %f %f %f\n", cx-h, cy+h, cz-h) // top l
	}

	fmt.Fprintf(file, "\n")

	// ini urutan rotasiny hrs pas, intiny buat facesny
	v := 1
	for i := 0; i < len(leaves); i++ {
		fmt.Fprintf(file, "f %d %d %d %d\n", v, v+1, v+2, v+3)   // front
		fmt.Fprintf(file, "f %d %d %d %d\n", v+5, v+4, v+7, v+6) // back
		fmt.Fprintf(file, "f %d %d %d %d\n", v+3, v+2, v+6, v+7) // top
		fmt.Fprintf(file, "f %d %d %d %d\n", v+4, v+5, v+1, v)   // bot
		fmt.Fprintf(file, "f %d %d %d %d\n", v+1, v+5, v+6, v+2) // r
		fmt.Fprintf(file, "f %d %d %d %d\n", v+4, v, v+3, v+7)   // l

		v += 8
	}

	return nil
}
