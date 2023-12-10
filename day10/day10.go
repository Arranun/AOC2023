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

type CurrentPos struct {
	position     [2]int
	pastPosition [2]int
	dirIn        [2]int
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	startPos := findPipeSPosition(&lines)
	fmt.Println(startPos)
	for _, l := range lines {
		fmt.Println(l)
	}
	followPipe(startPos, &lines)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getDirIn(pipe rune) [2]int {
	tmp := map[rune][2]int{'L': {-1, 1}, 'J': {-1, -1}, '7': {1, -1}, 'F': {1, 1}}
	return tmp[pipe]
}

func followPipe(start [2]int, lines *[]string) {
	pipePoints := map[[2]int]rune{}
	pos := CurrentPos{start, start, getDirIn(rune((*lines)[start[0]][start[1]]))}
	pipePoints[pos.position] = rune((*lines)[pos.position[0]][pos.position[1]])
	pos.getNextPos(rune((*lines)[pos.position[0]][pos.position[1]]))
	var step int
	for pos.position != start {
		pipePoints[pos.position] = rune((*lines)[pos.position[0]][pos.position[1]])
		pos.getNextPos(rune((*lines)[pos.position[0]][pos.position[1]]))
		step++
	}
	fmt.Println(step/2 + 1)
	tempLines := []string{}
	for i := 0; i < len(*lines); i++ {
		tempLine := ""
		for j := 0; j < len((*lines)[i]); j++ {
			val, ok := pipePoints[[2]int{i, j}]
			if ok {
				tempLine += string(val)
			} else {
				tempLine += "."
			}
		}
		tempLines = append(tempLines, tempLine)
	}
	for _, l := range tempLines {
		fmt.Println(l)
	}
	tempLines = removeTopBottom(tempLines)
	tempLines = removeLeftRight(tempLines)
	fmt.Println()
	for _, l := range tempLines {
		fmt.Println(l)
	}
}

func removeLeftRight(tempLines []string) []string {
	constainsPipe := func(i int) bool {
		for j, _ := range tempLines {
			if tempLines[j][i] != '.' {
				return true
			}
		}
		return false
	}
	right := 0
	for {
		if !constainsPipe(right) {
			right++
		} else {
			break
		}
	}
	left := len(tempLines[0]) - 1
	for {
		if !constainsPipe(left) {
			left--
		} else {
			break
		}
	}
	for i, v := range tempLines {
		tempLines[i] = v[right : left+1]
	}
	return tempLines
}

func removeTopBottom(tempLines []string) []string {
	constainsPipe := func(rune2 rune) bool {
		return rune2 != '.'
	}
	tmptmpLines := []string{}
	for _, l := range tempLines {
		if strings.ContainsFunc(l, constainsPipe) {
			tmptmpLines = append(tmptmpLines, l)
		}
	}
	tempLines = tmptmpLines
	return tempLines
}

func findPipeSPosition(lines *[]string) [2]int {
	var startPos [2]int
	for i, _ := range *lines {
		for j, _ := range (*lines)[i] {
			if (*lines)[i][j] == 'S' {
				startPos = [2]int{i, j}
				break
			}
		}
	}
	x := startPos[0]
	y := startPos[1]
	directions := [4][2]int{{x - 1, y}, {x, y + 1}, {x + 1, y}, {x, y - 1}}
	possiblePipes := [4][]string{{"|", "7", "F"}, {"-", "J", "7"}, {"|", "L", "J"}, {"-", "L", "F"}}
	var possibleDirections [2]int
	pdc := 0
	for i, d := range directions {
		if len(*lines) > d[0] && d[0] > -1 && d[1] > -1 && len((*lines)[d[0]]) > d[1] {
			if slices.Contains(possiblePipes[i], string((*lines)[d[0]][d[1]])) {
				possibleDirections[pdc] = i
				pdc++
			}
		}
	}
	temp := map[[2]int]string{{0, 1}: "L", {0, 2}: "|", {1, 3}: "-", {0, 3}: "J", {2, 3}: "7", {1, 2}: "F"}
	(*lines)[x] = strings.Replace((*lines)[x], "S", temp[possibleDirections], -1)
	return startPos
}

func (c *CurrentPos) getNextPos(pipe rune) {
	//| is a vertical pipe connecting north and south.
	//- is a horizontal pipe connecting east and west.
	//L is a 90-degree bend connecting north and east.
	//J is a 90-degree bend connecting north and west.
	//7 is a 90-degree bend connecting south and west.
	//F is a 90-degree bend connecting south and east.
	//. is ground; there is no pipe in this tile.
	switch pipe {
	case '|':
		c.chooseNextPositionBasedOnLastPosition([][2]int{{1, 0}, {-1, 0}})
	case '-':
		c.chooseNextPositionBasedOnLastPosition([][2]int{{0, -1}, {0, 1}})
	case 'L':
		c.chooseNextPositionBasedOnLastPosition([][2]int{{-1, 0}, {0, 1}})
	case 'J':
		c.chooseNextPositionBasedOnLastPosition([][2]int{{-1, 0}, {0, -1}})
	case '7':
		c.chooseNextPositionBasedOnLastPosition([][2]int{{1, 0}, {0, -1}})
	case 'F':
		c.chooseNextPositionBasedOnLastPosition([][2]int{{1, 0}, {0, 1}})
	}
}

func (c *CurrentPos) chooseNextPositionBasedOnLastPosition(possibleDir [][2]int) {
	possiblePos := make([][2]int, len(possibleDir))
	for i, d := range possibleDir {
		possiblePos[i] = [2]int{c.position[0] + d[0], c.position[1] + d[1]}
	}
	for _, p := range possiblePos {
		if p != c.pastPosition {
			(*c).pastPosition = (*c).position
			(*c).position = p
			break
		}
	}
}

func (c *CurrentPos) changeDirectionIn(dir [2]int) {
	switch dir {
	case [2]int{0, 1}:
		(*c).dirIn[1] = 0
		break
	case [2]int{0, -1}:
		(*c).dirIn[1] = 0
		break
	}

}
