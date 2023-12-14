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

type RowPart2 struct {
	groups string
	backup []int
}

type ActivPath struct {
	springIndex   int
	nextBackup    int
	currentBackup int
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	part2(lines)
	part2Lines := make([]string, len(lines))
	for i, r := range lines {
		fields := strings.Fields(r)
		part2Lines[i] = strings.Repeat(fields[0]+"?", 5)[:5*(len(fields[0])+1)-1] + " " + strings.Repeat(fields[1]+",", 5)[:5*(len(fields[1])+1)-1]
	}
	part2(part2Lines)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part2(lines []string) {
	rows := make([]RowPart2, len(lines))
	for i, l := range lines {
		split := strings.Fields(l)
		rows[i].groups = split[0]
		rows[i].backup = helper.StringSliceToIntSlice(strings.Split(split[1], ","))
	}
	sum := 0
	for _, r := range rows {
		possiblePaths := r.getPossiblePathsPart2()
		sum += possiblePaths
	}
	fmt.Println(sum)

}

func (row RowPart2) getPossiblePathsPart2() int {
	activePaths := map[ActivPath]int{}
	activePaths[ActivPath{0, 0, -1}] = 1
	possiblePaths := 0
	for len(activePaths) > 0 {
		var currentPath ActivPath
		var currentAmount int
		for currentPath, currentAmount = range activePaths {
			break
		}
		delete(activePaths, currentPath)
		if currentPath.currentBackup < 1 && currentPath.nextBackup == len(row.backup) && currentPath.springIndex == len(row.groups) {
			possiblePaths += currentAmount
		} else if row.enoughRemainingBroken(currentPath) {
			for _, r := range currentPath.getPossibleNextSteps(row) {
				tmpActivePath := currentPath
				if tmpActivePath.nextStep(r, row) {
					activePaths[tmpActivePath] += currentAmount
				}
			}
		}
	}
	return possiblePaths
}

func (row RowPart2) enoughRemainingBroken(currentPath ActivPath) bool {
	return currentPath.currentBackup+helper.Sum(row.backup[currentPath.nextBackup:]) <= strings.Count(row.groups[currentPath.springIndex:], "?")+strings.Count(row.groups[currentPath.springIndex:], "#")
}

func (aP *ActivPath) getPossibleNextSteps(row RowPart2) []rune {
	if aP.springIndex > len(row.groups)-1 {
		return []rune{}
	}
	nextGroupValue := row.groups[aP.springIndex]
	switch nextGroupValue {
	case '.':
		return []rune{rune(nextGroupValue)}
	case '#':
		return []rune{rune(nextGroupValue)}
	}
	return []rune{'#', '.'}
}

func (aP *ActivPath) nextStep(nextSpring rune, row RowPart2) bool {
	if aP.currentBackup > 0 && nextSpring == '#' {
		aP.currentBackup -= 1
		aP.springIndex++
		return true
	}
	if aP.currentBackup == 0 && nextSpring == '.' {
		aP.currentBackup -= 1
		aP.springIndex++
		return true
	}
	if aP.currentBackup == -1 && nextSpring == '.' {
		aP.springIndex++
		return true
	}
	if aP.nextBackup < len(row.backup) && aP.currentBackup == -1 && nextSpring == '#' {
		aP.currentBackup = row.backup[aP.nextBackup] - 1
		aP.nextBackup++
		aP.springIndex++
		return true
	}
	return false
}

func test(lines []string) {
	expection := make([]int, len(lines))
	for i, l := range lines {
		fields := strings.Fields(l)
		expection[i] = helper.RemoveError(strconv.Atoi(fields[2]))
	}
	rows := make([]RowPart2, len(lines))
	for i, l := range lines {
		split := strings.Fields(l)
		rows[i].groups = split[0]
		rows[i].backup = helper.StringSliceToIntSlice(strings.Split(split[1], ","))
	}
	sum := 0
	for i, r := range rows {
		possiblePaths := r.getPossiblePathsPart2()
		if !(possiblePaths == expection[i]) {
			fmt.Printf("%s: %d \n", lines[i], possiblePaths)
			possiblePaths = r.getPossiblePathsPart2()
		}
		sum += possiblePaths
	}
	fmt.Println(sum)
}
