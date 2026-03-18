package main

import (
	"fmt"
	"tucil/src/parser"
)

func main() {

    // Test Parsing
    vertices, faces := parser.Parse("tc/cow.obj")

    fmt.Println("==== VERTICES ====")
    for i, row := range vertices {
        fmt.Print(i, " : ")
        for j, col := range row {
            fmt.Printf("%.6f", col)
            if j != 2 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }

    fmt.Println("\n==== FACES ====")
    for i, row := range faces {
        fmt.Print(i, " : ")
        for j, col := range row {
            fmt.Printf("%d", col)
            if j != 2 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}