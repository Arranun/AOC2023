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
	pattern := [][]string{}
	currentPattern := []string{}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			pattern = append(pattern, currentPattern)
			currentPattern = []string{}
		} else {
			currentPattern = append(currentPattern, lines[i])
		}
	}
	pattern = append(pattern, currentPattern)
	sumPart1 := 0
	sumPart2 := 0
	for _, p := range pattern {
		sumPart1 = step(p, sumPart1, 0)
		sumPart2 = step(p, sumPart2, 1)
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func step(p []string, sumPart1 int, fix int) int {
	rows := checkRows(p, fix)
	columns := checkColumns(p, fix)
	if rows > -1 {
		sumPart1 += (rows + 1) * 100
	}
	if columns > -1 {
		sumPart1 += columns + 1
	}
	return sumPart1
}

func checkRows(lines []string, fix int) int {
	var allowedBreak int
	var check bool
	for i := 0; i < len(lines)-1; i++ {
		allowedBreak = fix
		allowedBreak, check = checkRow(lines, i, i+1, allowedBreak)
		if check {
			tempI := i - 1
			for tempI != -1 && i+1+(i-tempI) < len(lines) {
				allowedBreak, check = checkRow(lines, tempI, i+1+(i-tempI), allowedBreak)
				if check {
					tempI--
				} else {
					break
				}
			}
			if (tempI == -1 || i+1+(i-tempI) >= len(lines)) && allowedBreak == 0 {
				return i
			}
		}
	}
	return -1
}

func checkRow(lines []string, i int, i2 int, fix int) (int, bool) {
	allowedBreaks := fix
	for j, _ := range lines[i] {
		if lines[i][j] != lines[i2][j] {
			allowedBreaks--
			if allowedBreaks < 0 {
				return allowedBreaks, false
			}
		}
	}
	return allowedBreaks, true
}

func checkColumns(lines []string, fix int) int {
	var allowedBreak int
	var check bool
	for i := 0; i < len(lines[0])-1; i++ {
		allowedBreak = fix
		allowedBreak, check = checkColumn(lines, i, i+1, allowedBreak)
		if check {
			tempI := i - 1
			for tempI != -1 && i+1+(i-tempI) < len(lines[0]) {
				allowedBreak, check = checkColumn(lines, tempI, i+1+(i-tempI), allowedBreak)
				if check {
					tempI--
				} else {
					break
				}
			}
			if (tempI == -1 || i+1+(i-tempI) >= len(lines[0])) && allowedBreak == 0 {
				return i
			}
		}
	}
	return -1
}

func checkColumn(lines []string, i int, i2 int, fix int) (int, bool) {
	allowedBreaks := fix
	for j := 0; j < len(lines); j++ {
		if lines[j][i] != lines[j][i2] {
			allowedBreaks--
			if allowedBreaks < 0 {
				return allowedBreaks, false
			}
		}
	}
	return allowedBreaks, true
}
