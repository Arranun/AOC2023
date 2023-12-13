package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Row struct {
	groups []string
	backup []int
}

type Path struct {
	currentGroupIndex int
	remainBackup      []int
	amount            int
}

type GroupVariants struct {
	backup []int
	amount int
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	testBool := len(args) > 1 && args[1] == "test"
	start := time.Now()
	if testBool {
		test(lines)
	} else {
		part1(lines)
	}

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part2(lines []string) {
	newLines := make([]string, len(lines))
	for i, l := range lines {
		split := strings.Split(l, " ")
		newLines[i] = strings.Repeat(split[0]+"?", 5)[:5*(len(split[0])+1)-1] + " " + strings.Repeat(split[1]+",", 5)[:5*(len(split[1])+1)-1]
		fmt.Println(newLines[i])
	}
	part1(newLines)
}

func part1(lines []string) {
	rows := getRows(lines)
	groupVariantMap := getGroupVariants(rows)
	sum := 0
	for _, r := range rows {
		possiblePaths := getPossibleVariations(r, &groupVariantMap)
		sum += possiblePaths
	}
	fmt.Println(sum)
}

func getRows(lines []string) []Row {
	rows := make([]Row, len(lines))
	for i, line := range lines {
		split := strings.Fields(line)
		rows[i].groups = strings.Fields(strings.ReplaceAll(split[0], ".", " "))
		rows[i].backup = helper.StringSliceToIntSlice(strings.Split(split[1], ","))
	}
	return rows
}

func getGroupVariants(rows []Row) map[string][]GroupVariants {
	groupVariantMap := map[string][]GroupVariants{}
	for _, r := range rows {
		for _, g := range r.groups {
			_, ok := groupVariantMap[g]
			if !ok && len(r.groups) > 1 {
				groupVariantMap[g] = getVariationsOfGroup(g)
			}

		}
	}
	return groupVariantMap
}

func test(lines []string) {
	rowLines := make([]string, len(lines))
	expection := make([]int, len(lines))
	for i, l := range lines {
		fields := strings.Fields(l)
		rowLines[i] = fields[0] + " " + fields[1]
		expection[i] = helper.RemoveError(strconv.Atoi(fields[2]))
	}
	rows := getRows(rowLines)
	groupVariantMap := getGroupVariants(rows)
	sum := 0
	for i, r := range rows {
		possiblePaths := getPossibleVariations(r, &groupVariantMap)
		sum += possiblePaths
		if !(possiblePaths == expection[i]) {
			fmt.Printf("%s: %d \n", lines[i], possiblePaths)
			possiblePaths = getPossibleVariations(r, &groupVariantMap)
		}
	}
	fmt.Println(sum)
}

func getPossibleVariations(row Row, groupVariantMap *map[string][]GroupVariants) int {
	activePaths := []Path{{0, row.backup, 1}}
	possiblePaths := []Path{}
	for len(activePaths) > 0 {
		currentPath := activePaths[0]
		activePaths = activePaths[1:]
		if currentPath.currentGroupIndex == len(row.groups) && len(currentPath.remainBackup) == 0 {
			possiblePaths = append(possiblePaths, currentPath)
		} else if !(currentPath.currentGroupIndex == len(row.groups)) {
			groupVars := (*groupVariantMap)[row.groups[currentPath.currentGroupIndex]]
			for _, gv := range groupVars {
				possible, _ := groupVarPossible(gv.backup, currentPath.remainBackup)
				if possible {
					zeroesInGv := 0
					for _, v := range gv.backup {
						if v == 0 {
							zeroesInGv++
						}
					}
					newActivePath := Path{}
					newActivePath.currentGroupIndex = currentPath.currentGroupIndex + 1
					newActivePath.remainBackup = currentPath.remainBackup[len(gv.backup)-zeroesInGv:]
					newActivePath.amount = currentPath.amount * gv.amount
					activePaths = append(activePaths, newActivePath)
				}
			}
		}
	}
	sum := 0
	for _, p := range possiblePaths {
		sum += p.amount
	}
	return sum
}

func groupVarPossible(groupVar []int, remainingBackup []int) (bool, []int) {
	if remainingBackup[0] == 0 {
		if groupVar[0] != 0 {
			return false, remainingBackup
		}
		remainingBackup = remainingBackup[1:]
		groupVar = groupVar[1:]
	}

	if len(groupVar) > len(remainingBackup) {
		return false, remainingBackup
	}
	index := 0
	for _, gv := range groupVar {
		if gv != 0 && gv != remainingBackup[index] {
			return false, remainingBackup
		}
		index++
	}
	return true, remainingBackup[index:]
}

func getVariationsOfGroup(group string) []GroupVariants {
	unknownIndex := []int{}
	for i, c := range group {
		if c == '?' {
			unknownIndex = append(unknownIndex, i)
		}
	}
	variantsLen := int(math.Pow(2, float64(len(unknownIndex))))
	groupVariants := []GroupVariants{}
	groupVariantsMap := map[string]int{}
	for i := 0; i < variantsLen; i++ {
		binary := fmt.Sprintf("%0*b", len(unknownIndex), i)
		bv := make([]bool, len(unknownIndex))
		for j, b := range binary {
			if b == '1' {
				bv[j] = true
			}
		}
		backupFromGroupVariant := getBackupFromGroupVariant(group, unknownIndex, bv)
		groupVariantsMap[strings.Trim(strings.Join(strings.Fields(fmt.Sprint(backupFromGroupVariant)), ","), "[]")]++
	}
	for k, v := range groupVariantsMap {
		backup := strings.Split(k, ",")
		tempVariant := make([]int, len(backup))
		for i, n := range backup {
			tempVariant[i] = helper.RemoveError(strconv.Atoi(n))
		}
		groupVariants = append(groupVariants, GroupVariants{tempVariant, v})
	}

	return groupVariants
}

func getBackupFromGroupVariant(group string, unknownIndex []int, binary []bool) []int {
	backup := []int{}
	currentBackup := 0
	unknownIndexIndex := 0
	binaryIndex := 0
	for i, r := range group {
		if unknownIndexIndex < len(unknownIndex) && unknownIndex[unknownIndexIndex] == i {
			if binary[binaryIndex] {
				r = '#'
			} else {
				r = '.'
			}
			binaryIndex++
			unknownIndexIndex++
		}
		if r == '#' {
			currentBackup++
		}
		if r == '.' {
			backup = append(backup, currentBackup)
			currentBackup = 0
		}
	}
	backup = append(backup, currentBackup)
	lastN := -1
	tmpBackup := []int{}
	for _, n := range backup {
		if n != 0 {
			tmpBackup = append(tmpBackup, n)
		}
		if n == 0 && lastN != 0 {
			tmpBackup = append(tmpBackup, n)
		}
		lastN = n
	}

	return tmpBackup
}

func getBackupFromGroup(group string) []int {
	backup := []int{0}
	currentBackup := 0
	for _, r := range group {
		if r == '#' {
			currentBackup++
		}
		if r == '.' && currentBackup > 0 {
			backup = append(backup, currentBackup)
			currentBackup = 0
		}
	}
	return backup
}
