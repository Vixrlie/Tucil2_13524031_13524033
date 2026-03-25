package main

import (
	"fmt"
	"log"
	"tucil/src/octree"
	"tucil/src/parser"
	"tucil/src/point"
	"tucil/src/voxelizer"
)

func main() {
	// Parse
	vertices, faces := parser.Parse("tc/bunny.obj")

	var arrPoints []point.Point
	arrPoints = point.ToPoint(vertices)

	var arrFaces []point.Face
	arrFaces = point.ToFace(arrPoints, faces)

	targetSize := 0.00390625

	fmt.Println("Voxelizing...")

	rootNode := voxelizer.StartVoxelize(arrPoints, arrFaces, targetSize)

	var leaves []*octree.OctreeNode
	voxelizer.GetLeaves(rootNode, &leaves)

	fmt.Printf("Generated %d solid voxels!\n", len(leaves))

	outputFile := "test/placeholder.obj"
	errr := voxelizer.ExportToOBJ(outputFile, leaves)
	if errr != nil {
		log.Fatalf("Failed to save voxels: %v", errr)
	}
}
