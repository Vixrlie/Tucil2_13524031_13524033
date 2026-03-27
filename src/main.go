package main

import (
	"fmt"
	"log"
	"math"
	"time"
	"tucil/src/octree"
	"tucil/src/parser"
	"tucil/src/point"
	"tucil/src/voxelizer"
)

func main() {
	var filename string
	var targetSize float64

	// Input
	fmt.Println("\n=== MAMBA VOXELIZER ===")
	fmt.Println("Please enter the filename you desire")
	fmt.Println("Hint : file located in tc/")
	fmt.Print(">> ")
	fmt.Scanf("%s\n", &filename)

	targetSize = math.Pow(2, targetSize)

	// Parse
	vertices, faces := parser.Parse("tc/" + filename)

	var arrPoints []point.Point
	arrPoints = point.ToPoint(vertices)

	var arrFaces []point.Face
	arrFaces = point.ToFace(arrPoints, faces)

	fmt.Println("\n=== STARTS VOXELIZING ===")

	rootNode := voxelizer.StartVoxelize(arrPoints, arrFaces)

	start_time := time.Now()

	fmt.Println("\n=== OUTPUT ===")
	var leaves []*octree.OctreeNode
	voxelizer.GetLeaves(rootNode, &leaves)

	fmt.Println("--- Voxels Generated ---")
	fmt.Printf("vx : %v\n\n", len(leaves))

	// print vertex & face
	outputFile := "test/v_" + filename
	errr := voxelizer.ExportToOBJ(outputFile, leaves)
	if errr != nil {
		log.Fatalf("Failed to save voxels: %v", errr)
	}

	// print nodes detail
	voxelizer.PrintDepthDetails()

	fmt.Printf("--- File saved in %s ---\n\n", outputFile)
	elapsed := time.Since(start_time)
	log.Printf("The process took %s", elapsed)
}
