package main

import (
	"fmt"
	"os"
	"tucil/src/octree"
	"tucil/src/parser"
	"tucil/src/point"
	"tucil/src/wrapper"
)

func main() {
	output_file, err := os.Create("output.obj")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Test Parsing
	vertices, _ := parser.Parse("test/diamond.obj") //Diganti dengan path ke testcase nanti

	var arrPoints []point.Point
	arrPoints = point.ToPoint(vertices)

	var oct *octree.OctreeNode = nil
	arrWrapper, oct := wrapper.WrapInBox(arrPoints)
	if oct == nil {
		fmt.Println("kosong")
	}

	point.PrintPoints(arrWrapper)
	point.FPrintPoints(output_file, arrWrapper)

	// fmt.Fprintf(output_file, "\n\nf 1 2 3\nf 2 3 4\nf 4 3 7\nf 4 7 8\nf 7 8 6\nf 7 6 5\nf 5 1 2\nf 5 2 6\nf 1 3 5\nf 3 5 7\nf 2 4 6\nf 4 6 8\n")

	fmt.Fprintf(output_file, "\n\nf 1 3 7 5\nf 2 4 8 6\nf 1 2 4 3\nf 5 6 8 7\nf 1 2 6 5\nf 3 4 8 7\n")

	// fmt.Println("==== VERTICES ====")
	// fmt.Fprintln(output_file, "# ==== VERTICES ====")
	// for i, row := range vertices {
	// 	fmt.Print(i, " : ")
	// 	fmt.Fprintf(output_file, "v ")
	// 	for j, col := range row {
	// 		fmt.Printf("%.6f", col)
	// 		fmt.Fprintf(output_file, "%.6f", col)
	// 		if j != 2 {
	// 			fmt.Print(" ")
	// 			fmt.Fprint(output_file, " ")
	// 		}
	// 	}
	// 	fmt.Println()
	// 	fmt.Fprintln(output_file)
	// }

	// fmt.Println("\n==== FACES ====")
	// fmt.Fprintln(output_file, "\n\n# ==== FACES ====")

	// for i, row := range faces {
	// 	fmt.Print(i, " : ")
	// 	fmt.Fprintf(output_file, "f ")
	// 	for j, col := range row {
	// 		fmt.Printf("%d", col)
	// 		fmt.Fprintf(output_file, "%d", col)
	// 		if j != 2 {
	// 			fmt.Print(" ")
	// 			fmt.Fprint(output_file, " ")
	// 		}
	// 	}
	// 	fmt.Println()
	// 	fmt.Fprintln(output_file)
	// }

	defer output_file.Close()

}
