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
	part2Commands := make([][]string, len(lines))
	directions := [4]string{"R", "D", "L", "U"}
	for i, l := range commands {
		r, _ := strconv.ParseInt(l[2][2:len(l[2])-2], 16, 64)
		d, _ := strconv.ParseInt(l[2][len(l[2])-2:len(l[2])-1], 16, 64)
		part2Commands[i] = []string{directions[d], strconv.FormatInt(r, 10)}
	}
	part1(commands)
	part1(part2Commands)
	start := time.Now()
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part1(commands [][]string) {
	point := helper.Point{0, 0, getDirection(commands[0]), 1, [][4]int{}}
	for _, c := range commands {
		length := helper.RemoveError(strconv.Atoi(c[1]))
		point.Direction = getDirection(c)
		point.StepLong(length)
	}
	var area int
	area = shoelace(point)
	fmt.Println()
	fmt.Println(area - point.PathLength/2 + point.PathLength)
}

func shoelace(point helper.Point) int {
	sum := 0
	points := make([][2]int, len(point.History)+1)
	for i, hpoint := range point.History {
		points[i] = [2]int{hpoint[0], hpoint[1]}
	}
	points[len(points)-1] = [2]int{point.History[0][1], point.History[0][0]}
	for i := 0; i < len(points)-1; i++ {
		currentPoint := points[i]
		nextPoint := points[i+1]
		sum += currentPoint[0] * nextPoint[1]
		sum -= currentPoint[1] * nextPoint[0]
	}
	return -1 * sum / 2
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
