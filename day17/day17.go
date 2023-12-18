package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Map struct {
	layout            [][]int
	fastestConnection map[[4]int]Point
	activePoints      []Point
}

type Point struct {
	basePoint        helper.Point
	remainingForward int
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	baseStartPoint := helper.Point{0, 0, [2]int{0, 1}, 0, [][4]int{}}
	points := map[[4]int]Point{[4]int{0, 0, 0, 1}: {baseStartPoint, 0}}
	layout := make([][]int, len(lines))
	for i, l := range lines {
		layout[i] = helper.StringSliceToIntSlice(strings.Split(l, ""))
	}
	m := Map{layout, points, []Point{{baseStartPoint, 0}}}
	for len(m.activePoints) > 0 {
		m.step()
	}
	fmt.Println(m.getMinimumValueForPos(len(m.layout)-1, len(m.layout[0])-1))
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func (m *Map) stopCondition() bool {
	g := [2]int{len(m.layout) - 1, len(m.layout[0]) - 1}
	for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		if m.fastestConnection[[4]int{g[0], g[1], d[0], d[1]}].basePoint.PathLength != 0 {
			point := m.fastestConnection[[4]int{g[0], g[1], d[0], d[1]}]
			printHistory(point.basePoint, m.layout)
			fmt.Println(point.basePoint.PathLength)
			return true
		}
	}
	return false
}

func printHistory(p helper.Point, layout [][]int) {
	lines := make([][]string, len(layout))
	for i, l := range layout {
		lines[i] = make([]string, len(l))
		for j, n := range l {
			lines[i][j] = strconv.Itoa(n)
		}
	}
	for _, pos := range p.History {
		if pos[2] == 0 && pos[3] == 1 {
			lines[pos[0]][pos[1]] = ">"
		}
		if pos[2] == 0 && pos[3] == -1 {
			lines[pos[0]][pos[1]] = "<"
		}
		if pos[2] == 1 && pos[3] == 0 {
			lines[pos[0]][pos[1]] = "v"
		}
		if pos[2] == -1 && pos[3] == 0 {
			lines[pos[0]][pos[1]] = "^"
		}
	}
	for _, l := range lines {
		fmt.Println(l)
	}
}

func (m *Map) step() {
	nextPoint := m.getNextPoint()
	currentPoints := getPointsWithDirectionChange(nextPoint)
	for _, currentPoint := range currentPoints {
		i := 0
		for i < 3 {
			currentPoint.basePoint.History = append([][4]int(nil), currentPoint.basePoint.History...)
			currentPoint.Step()
			if currentPoint.remainingForward < 0 ||
				currentPoint.basePoint.FromTop < 0 ||
				currentPoint.basePoint.FromLeft < 0 ||
				currentPoint.basePoint.FromTop >= len(m.layout) ||
				currentPoint.basePoint.FromLeft >= len(m.layout[0]) {
				break
			}
			currentPoint.basePoint.PathLength += m.layout[currentPoint.basePoint.FromTop][currentPoint.basePoint.FromLeft] - 1
			currentFastestConnection := m.fastestConnection[currentPoint.basePoint.GetPosAndDir()].basePoint.PathLength
			if currentFastestConnection == 0 || currentFastestConnection > currentPoint.basePoint.PathLength {
				m.fastestConnection[currentPoint.basePoint.GetPosAndDir()] = currentPoint
				if compare(currentPoint.basePoint.History, [][4]int{{0, 0, 1, 0}, {1, 0, 0, 1}, {1, 1, 0, 1}, {1, 2, 0, 1}, {0, 2, 0, 1}, {0, 3, 0, 1}, {1, 3, 1, 0}, {2, 3, 1, 0}}) {
					if currentPoint.basePoint.FromTop == 1 && currentPoint.basePoint.FromLeft == 5 {
						fmt.Println("Here")
					}
				}
				m.activePoints = append(m.activePoints, currentPoint)
			}
			i++
		}
	}
}

func compare(a, b [][4]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, a1 := range a {
		if a1 != b[i] {
			return false
		}
	}
	return true
}

func (m *Map) getNextPoint() Point {
	minimum := m.activePoints[0].basePoint.PathLength
	point := m.activePoints[0]
	pIndex := 0
	for i, p := range m.activePoints {
		if minimum > p.basePoint.PathLength {
			minimum = p.basePoint.PathLength
			point = p
			pIndex = i
		}
	}
	m.activePoints = slices.Delete(m.activePoints, pIndex, pIndex+1)
	return point
}

func (m *Map) print() {
	lines := make([][]int, len(m.layout))
	for i, _ := range m.layout {
		lines[i] = make([]int, len(m.layout[0]))
		for j, _ := range m.layout[0] {
			lines[i][j] = m.getMinimumValueForPos(i, j)
		}
		fmt.Println(lines[i])
	}
}

func (m *Map) getMinimumValueForPos(fromTop, fromLeft int) int {
	minimum := m.fastestConnection[[4]int{fromTop, fromLeft, 0, 1}].basePoint.PathLength
	minimumPoint := m.fastestConnection[[4]int{fromTop, fromLeft, 0, 1}]
	for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		if m.fastestConnection[[4]int{fromTop, fromLeft, d[0], d[1]}].basePoint.PathLength != 0 &&
			m.fastestConnection[[4]int{fromTop, fromLeft, d[0], d[1]}].basePoint.PathLength < minimum {
			minimumPoint = m.fastestConnection[[4]int{fromTop, fromLeft, d[0], d[1]}]
			minimum = minimumPoint.basePoint.PathLength
		}
	}
	printHistory(minimumPoint.basePoint, m.layout)
	return minimum
}

func getPointsWithDirectionChange(p Point) []Point {
	points := []Point{p, p, p}
	switch p.basePoint.Direction {
	case [2]int{0, 1}:
		fallthrough
	case [2]int{0, -1}:
		points[0].basePoint.Direction = [2]int{-1, 0}
		points[0].remainingForward = 3
		points[1].basePoint.Direction = [2]int{1, 0}
		points[1].remainingForward = 3
	case [2]int{1, 0}:
		fallthrough
	case [2]int{-1, 0}:
		points[0].basePoint.Direction = [2]int{0, 1}
		points[0].remainingForward = 3
		points[1].basePoint.Direction = [2]int{0, -1}
		points[1].remainingForward = 3
	}
	return points
}

func (p *Point) Step() {
	p.basePoint.Step()
	p.remainingForward--
}
