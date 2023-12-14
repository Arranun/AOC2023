package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	rowLoad := make([]int, len(lines))
	tiltNorth(lines, rowLoad)
	var sum int
	for i, r := range rowLoad {
		sum += r * (len(rowLoad) - i)
	}
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func tiltNorth(lines []string, rowLoad []int) {
	for i := 0; i < len(lines[0]); i++ {
		currentGroundRow := -1
		for j := 0; j < len(lines); j++ {
			switch lines[j][i] {
			case '#':
				currentGroundRow = j
			case 'O':
				rowLoad[currentGroundRow+1]++
				currentGroundRow++
			}
		}
	}
}
