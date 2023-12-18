package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	commands := make([][]string, len(lines))
	for i, l := range lines {
		commands[i] = strings.Fields(l)
	}
	point := helper.Point{0, 0, getDirection(commands[0]), 1, [][4]int{}}
	for _, c := range commands {
		length := helper.RemoveError(strconv.Atoi(c[1]))
		point.Direction = getDirection(c)
		for i := 0; i < length; i++ {
			point.Step()
		}
	}
	layout := printHistory(point)
	var area int
	layout, area = fillLoop(layout)
	fmt.Println()
	for _, l := range layout {
		fmt.Println(l)
	}
	fmt.Println(area + len(point.History))
	start := time.Now()
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func printHistory(p helper.Point) [][]string {
	highestFromTop := 0
	lowestFromTop := 0
	highestFromLeft := 0
	lowestFromLeft := 0
	for _, h := range p.History {
		if h[0] > highestFromTop {
			highestFromTop = h[0]
		}
		if h[0] < lowestFromTop {
			lowestFromTop = h[0]
		}
		if h[1] > highestFromLeft {
			highestFromLeft = h[1]
		}
		if h[1] < lowestFromLeft {
			lowestFromLeft = h[1]
		}
	}
	layout := make([][]string, highestFromTop-lowestFromTop+1)
	for i, _ := range layout {
		layout[i] = make([]string, highestFromLeft-lowestFromLeft+1)
		for j, _ := range layout[0] {
			layout[i][j] = "."
		}
	}
	lastDir := "^"
	for _, pos := range p.History {
		currentDir := ""
		if pos[2] == 0 && pos[3] == 1 {
			currentDir = ">"
		}
		if pos[2] == 0 && pos[3] == -1 {
			currentDir = "<"
		}
		if pos[2] == 1 && pos[3] == 0 {
			currentDir = "v"
		}
		if pos[2] == -1 && pos[3] == 0 {
			currentDir = "^"
		}
		if (lastDir == "v" || lastDir == "^") && (currentDir == "<" || currentDir == ">") {
			layout[pos[0]-lowestFromTop][pos[1]-lowestFromLeft] = lastDir
		} else {
			layout[pos[0]-lowestFromTop][pos[1]-lowestFromLeft] = currentDir
		}
		lastDir = currentDir
	}
	for _, l := range layout {
		fmt.Println(l)
	}
	return layout
}

func fillLoop(layout [][]string) ([][]string, int) {
	area := 0
	for i, l := range layout {
		lastUp := -1
		lastDown := -1
		for j, r := range l {
			if r == "^" {
				lastUp = j
			}
			if lastUp > -1 && r == "v" {
				lastDown = j
			}
			if lastUp > -1 && lastDown > -1 {
				for jj := lastUp + 1; jj < lastDown; jj++ {
					if layout[i][jj] == "." {
						layout[i][jj] = "#"
						area++
					}
				}
				lastUp = -1
				lastDown = -1
			}
		}
	}
	return layout, area
}

func getDirection(command []string) [2]int {
	switch command[0] {
	case "R":
		return [2]int{0, 1}
	case "L":
		return [2]int{0, -1}
	case "D":
		return [2]int{1, 0}
	case "U":
		return [2]int{-1, 0}
	default:
		return [2]int{0, 0}
	}
}
