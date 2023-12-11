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
	pointsPart1 := part1(lines)
	pointsPart2 := part2(lines, args, 1000000)
	getSum(pointsPart1)
	getSum(pointsPart2)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part2(lines []string, args []string, expansion int) [][2]int {
	lines = helper.ReadTextFile(args[0])
	emptyColumn := getEmptyColumns(lines)
	emptyRows := getEmptyRows(lines)
	unexpandedPoints := getPoints(lines)
	expandPoints := [][2]int{}
	for _, p := range unexpandedPoints {
		expandPoints = append(expandPoints, unexpandedPointToExpandedPoint(p, emptyRows, emptyColumn, expansion-1))
	}
	return expandPoints
}

func unexpandedPointToExpandedPoint(point [2]int, emptyColumns []int, emptyRows []int, expansion int) [2]int {
	indexColumn := 0
	for indexColumn < len(emptyColumns) && point[0] > emptyColumns[indexColumn] {
		indexColumn++
	}
	indexRow := 0
	for indexRow < len(emptyRows) && point[1] > emptyRows[indexRow] {
		indexRow++
	}
	return [2]int{point[0] + indexColumn*expansion, point[1] + indexRow*expansion}
}

func part1(lines []string) [][2]int {
	lines = expandLeftRight(lines)
	lines = expandTopBottom(lines)
	points := getPoints(lines)
	return points
}

func getSum(points [][2]int) {
	sum := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			sum += helper.ManHattanDistance(p1, p2)
		}
	}
	fmt.Println(sum)
}

func getPoints(lines []string) [][2]int {
	points := [][2]int{}
	for i, l := range lines {
		for j, r := range l {
			if r == '#' {
				points = append(points, [2]int{i, j})
			}
		}
	}
	return points
}

func expandLeftRight(tempLines []string) []string {
	emptyColumns := getEmptyColumns(tempLines)
	for i, v := range tempLines {
		line := ""
		index := 0
		left := 0
		right := 0
		for index < len(emptyColumns) {
			left = right
			right = emptyColumns[index] + 1
			line += v[left:right] + "."
			index++
		}
		line += v[right:]
		tempLines[i] = line
	}
	return tempLines
}

func getEmptyColumns(tempLines []string) []int {
	containsGalaxy := func(i int) bool {
		for j, _ := range tempLines {
			if tempLines[j][i] != '.' {
				return true
			}
		}
		return false
	}
	emptyColumns := []int{}
	for i, _ := range tempLines {
		if !containsGalaxy(i) {
			emptyColumns = append(emptyColumns, i)
		}
	}
	return emptyColumns
}

func expandTopBottom(lines []string) []string {
	emptyRows := getEmptyRows(lines)
	tmpLines := []string{}
	index := 0
	for i, l := range lines {
		if index < len(emptyRows) && i == emptyRows[index] {
			index++
			tmpLines = append(tmpLines, l)
		}
		tmpLines = append(tmpLines, l)
	}
	return tmpLines
}

func getEmptyRows(lines []string) []int {
	containsGalaxy := func(rune2 rune) bool {
		return rune2 != '.'
	}
	emptyRows := []int{}
	for i, l := range lines {
		if !strings.ContainsFunc(l, containsGalaxy) {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}
