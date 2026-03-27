package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func HandleFile(path string) (res *os.File) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func Parse(path string) (res1 [][3]float32, res2 [][3]int) {
	file := HandleFile(path)
	defer file.Close()
	buf := bufio.NewScanner(file)
	for buf.Scan() {
		line := buf.Text()
		if strings.HasPrefix(line, "v ") {
			var a, b, c float32
			fmt.Sscanf(line, "v %f %f %f", &a, &b, &c)
			temp := [3]float32{a, b, c}
			res1 = append(res1, temp)
		} else if strings.HasPrefix(line, "f ") {
			token := strings.Fields(line)
			var temp [3]int
			for i := 0; i < 3; i++ {
				// Ambil bagian sebelum slash pertama
				a := strings.Split(token[i+1], "/")[0]
				val, err := strconv.Atoi(a)
				if err == nil {
					temp[i] = val
				}
			}
			res2 = append(res2, temp)
		}
	}
	if err := buf.Err(); err != nil {
		log.Fatal(err)
	}

	return res1, res2
}
