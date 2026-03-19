package main

import (
	"fmt"
	"os"
	"tucil/src/parser"
)

func main() {
	output_file, err := os.Create("output.obj")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Test Parsing
	vertices, faces := parser.Parse("test/diamond.obj") //Diganti dengan path ke testcase nanti

	fmt.Println("==== VERTICES ====")
	fmt.Fprintln(output_file, "# ==== VERTICES ====")
	for i, row := range vertices {
		fmt.Print(i, " : ")
		fmt.Fprintf(output_file, "v ")
		for j, col := range row {
			fmt.Printf("%.6f", col)
			fmt.Fprintf(output_file, "%.6f", col)
			if j != 2 {
				fmt.Print(" ")
				fmt.Fprint(output_file, " ")
			}
		}
		fmt.Println()
		fmt.Fprintln(output_file)
	}

	fmt.Println("\n==== FACES ====")
	fmt.Fprintln(output_file, "\n\n# ==== FACES ====")

	for i, row := range faces {
		fmt.Print(i, " : ")
		fmt.Fprintf(output_file, "f ")
		for j, col := range row {
			fmt.Printf("%d", col)
			fmt.Fprintf(output_file, "%d", col)
			if j != 2 {
				fmt.Print(" ")
				fmt.Fprint(output_file, " ")
			}
		}
		fmt.Println()
		fmt.Fprintln(output_file)
	}

	defer output_file.Close()

}
