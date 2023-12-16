package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Map struct {
	layout        [][]string
	beams         [][]string
	currentPoints []Point
	pastPoints    map[Point]bool
}

type Point struct {
	fromTop   int
	fromLeft  int
	direction [2]int
}

func main() {
	args := os.Args[1:]
	start := time.Now()
	layout := getLayout(args[0])
	sum := part1(layout, getLayout(args[0]), Point{0, -1, [2]int{0, 1}})
	fmt.Println(sum)
	entrances := getPossibleEntrances(layout)
	for _, e := range entrances {
		tmpSum := part1(layout, getLayout(args[0]), e)
		if tmpSum > sum {
			sum = tmpSum
		}
	}
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getLayout(file string) [][]string {
	lines := helper.ReadTextFile(file)
	layout := make([][]string, len(lines))
	for i, l := range lines {
		layout[i] = strings.Split(l, "")
	}
	return layout
}

func getPossibleEntrances(layout [][]string) []Point {
	entrances := []Point{}
	mapWidth := len(layout)
	mapHeight := len(layout[0])
	for i := 0; i < mapWidth; i++ {
		entrances = append(entrances, Point{-1, i, [2]int{1, 0}})
		entrances = append(entrances, Point{mapHeight, i, [2]int{-1, 0}})
	}
	for i := 0; i < mapHeight; i++ {
		entrances = append(entrances, Point{i, -1, [2]int{0, 1}})
		entrances = append(entrances, Point{i, mapWidth, [2]int{0, -1}})
	}
	return entrances
}

func part1(layout [][]string, beams [][]string, startPoint Point) int {
	m := Map{layout, beams, []Point{startPoint}, make(map[Point]bool)}
	for len(m.currentPoints) > 0 {
		m.step()
	}
	sum := 0
	for _, l := range m.beams {
		sum += strings.Count(strings.Join(l, ""), "#")
	}
	return sum
}

func (m *Map) step() {
	newPoints := []Point{}
	filteredPoints := []Point{}
	for _, p := range m.currentPoints {
		if !m.pastPoints[p] {
			filteredPoints = append(filteredPoints, p)
			m.pastPoints[p] = true
		}
	}
	for _, p := range filteredPoints {
		p.step(&m.layout, &m.beams)
		if p.fromTop > -1 && p.fromLeft > -1 && !(p.fromTop >= len(m.layout)) && !(p.fromLeft >= len(m.layout[0])) {
			newPoints = append(newPoints, p.changeDir(m.layout[p.fromTop][p.fromLeft])...)
			newPoints = append(newPoints, p)
		}
	}
	m.currentPoints = newPoints
}

func (p *Point) step(layout *[][]string, beams *[][]string) {
	stops := p.getStops()
	direction := p.direction
	p.fromTop += direction[0]
	p.fromLeft += direction[1]
	for p.fromLeft > -1 && p.fromTop > -1 && p.fromLeft < len((*layout)[0]) && p.fromTop < len(*layout) {
		(*beams)[p.fromTop][p.fromLeft] = "#"
		if slices.Contains(stops, (*layout)[p.fromTop][p.fromLeft]) {
			return
		}
		p.fromTop += direction[0]
		p.fromLeft += direction[1]
	}
}

func (p *Point) changeDir(mirror string) []Point {
	newPoints := []Point{}
	if p.direction == [2]int{1, 0} {
		switch mirror {
		case "/":
			p.direction = [2]int{0, -1}
			return newPoints
		case "\\":
			p.direction = [2]int{0, 1}
			return newPoints
		case "-":
			p.direction = [2]int{0, 1}
			return append(newPoints, Point{p.fromTop, p.fromLeft, [2]int{0, -1}})
		}
	}
	if p.direction == [2]int{-1, 0} {
		switch mirror {
		case "/":
			p.direction = [2]int{0, 1}
			return newPoints
		case "\\":
			p.direction = [2]int{0, -1}
			return newPoints
		case "-":
			p.direction = [2]int{0, 1}
			return append(newPoints, Point{p.fromTop, p.fromLeft, [2]int{0, -1}})
		}
	}
	if p.direction == [2]int{0, 1} {
		switch mirror {
		case "/":
			p.direction = [2]int{-1, 0}
			return newPoints
		case "\\":
			p.direction = [2]int{1, 0}
			return newPoints
		case "|":
			p.direction = [2]int{1, 0}
			return append(newPoints, Point{p.fromTop, p.fromLeft, [2]int{-1, 0}})
		}
	}
	if p.direction == [2]int{0, -1} {
		switch mirror {
		case "/":
			p.direction = [2]int{1, 0}
			return newPoints
		case "\\":
			p.direction = [2]int{-1, 0}
			return newPoints
		case "|":
			p.direction = [2]int{1, 0}
			return append(newPoints, Point{p.fromTop, p.fromLeft, [2]int{-1, 0}})
		}
	}
	return newPoints
}

func (p *Point) getStops() []string {
	stops := []string{"/", "\\", "-", "|"}
	switch p.direction {
	case [2]int{1, 0}:
		fallthrough
	case [2]int{-1, 0}:
		stops = slices.DeleteFunc(stops, func(stop string) bool { return stop == "|" })
	case [2]int{0, 1}:
		fallthrough
	case [2]int{0, -1}:
		stops = slices.DeleteFunc(stops, func(stop string) bool { return stop == "-" })
	}
	return stops
}
