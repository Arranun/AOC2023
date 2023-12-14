package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	scores := map[string]int{}
	var cycle int
	for i := 0; i < 1000; i++ {
		lines = step(lines)
		key := getKey(lines)
		if scores[key] > 0 {
			cycle = i - scores[key]
		} else {
			scores[key] = i
		}
	}
	lines = helper.ReadTextFile(args[0])
	for i := 0; i < (1000000000 % cycle); i++ {
		lines = step(lines)
	}
	sum := getScore(lines)
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func step(lines []string) []string {
	lines = tiltNorth(lines)
	lines = tiltWest(lines)
	lines = tiltSouth(lines)
	lines = tiltEast(lines)
	return lines
}

func getKey(lines []string) string {
	line := ""
	for _, l := range lines {
		line += l
	}
	return line
}

func getScore(lines []string) int {
	var sum int
	for i, l := range lines {
		sum += (len(lines) - i) * strings.Count(l, "O")
	}
	return sum
}

func compareLines(lines1, lines2 []string) bool {
	for i, l := range lines1 {
		if lines2[i] != l {
			return false
		}
	}
	return true
}

func tiltWest(lines []string) []string {
	tmpLines := make([]string, len(lines))
	for i, l := range lines {
		tmpLines[i] = l
	}
	for i := 0; i < len(lines); i++ {
		currentGround := -1
		for j := 0; j < len(lines[0]); j++ {
			switch lines[i][j] {
			case '#':
				currentGround = j
			case 'O':
				tmpLines[i] = tmpLines[i][:j] + "." + tmpLines[i][j+1:]
				tmpLines[i] = tmpLines[i][:currentGround+1] + "O" + tmpLines[i][currentGround+2:]
				currentGround++
			}
		}
	}
	return tmpLines
}

func tiltEast(lines []string) []string {
	tmpLines := make([]string, len(lines))
	for i, l := range lines {
		tmpLines[i] = l
	}
	for i := 0; i < len(lines); i++ {
		currentGround := len(lines[0])
		for j := len(lines[0]) - 1; j > -1; j-- {
			switch lines[i][j] {
			case '#':
				currentGround = j
			case 'O':
				tmpLines[i] = tmpLines[i][:j] + "." + tmpLines[i][j+1:]
				part1 := tmpLines[i][:currentGround-1]
				part2 := tmpLines[i][currentGround:]
				tmpLines[i] = part1 + "O" + part2
				currentGround--
			}
		}
	}
	return tmpLines
}

func tiltNorth(lines []string) []string {
	tmpLines := make([]string, len(lines))
	for i, l := range lines {
		tmpLines[i] = l
	}
	for i := 0; i < len(lines[0]); i++ {
		currentGroundRow := -1
		for j := 0; j < len(lines); j++ {
			switch lines[j][i] {
			case '#':
				currentGroundRow = j
			case 'O':
				tmpLines[j] = tmpLines[j][:i] + "." + tmpLines[j][i+1:]
				tmpLines[currentGroundRow+1] = tmpLines[currentGroundRow+1][:i] + "O" + tmpLines[currentGroundRow+1][i+1:]
				currentGroundRow++
			}
		}
	}
	return tmpLines
}

func tiltSouth(lines []string) []string {
	tmpLines := make([]string, len(lines))
	for i, l := range lines {
		tmpLines[i] = l
	}
	for i := 0; i < len(lines[0]); i++ {
		currentGroundRow := len(lines)
		for j := len(lines) - 1; j > -1; j-- {
			switch lines[j][i] {
			case '#':
				currentGroundRow = j
			case 'O':
				tmpLines[j] = tmpLines[j][:i] + "." + tmpLines[j][i+1:]
				tmpLines[currentGroundRow-1] = tmpLines[currentGroundRow-1][:i] + "O" + tmpLines[currentGroundRow-1][i+1:]
				currentGroundRow--
			}
		}
	}
	return tmpLines
}
