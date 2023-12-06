package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type mapElement struct {
	Destination int
	Source      int
	Range       int
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	seeds := helper.StringSliceToIntSlice(strings.Split(lines[0][strings.Index(lines[0], ": ")+2:], " "))
	maps := [][]mapElement{}
	activemap := []mapElement{}
	for i := 0; i < len(lines[3:]); i++ {
		l := lines[3:][i]
		if len(l) == 0 {
			i++
			maps = append(maps, activemap)
			activemap = []mapElement{}
		} else {
			destinationSourceRange := helper.StringSliceToIntSlice(strings.Split(l, " "))
			activemap = append(activemap, mapElement{destinationSourceRange[0], destinationSourceRange[1], destinationSourceRange[2]})
		}
	}
	maps = append(maps, activemap)
	//seedsPart2 := []int{}
	startRanges := [][2]int{}
	for i := 0; i < len(seeds); i += 2 {
		startRanges = append(startRanges, [2]int{seeds[i], seeds[i+1]})
		//for j := 0; j < seeds[i+1]; j++ {
		//	seedsPart2 = append(seedsPart2, seeds[i]+j)
		//}
	}
	//getLowestPossibleSoil(maps, seeds)
	//getLowestPossibleSoil(maps, seedsPart2)
	for _, m := range maps {
		startRanges = part2Step(startRanges, m)
	}
	var minVal = startRanges[0][0]
	for _, v := range startRanges {
		if minVal > v[0] {
			minVal = v[0]
		}
	}
	fmt.Println(minVal)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getLowestPossibleSoil(maps [][]mapElement, seeds []int) {
	fmt.Println(seeds)
	for _, m := range maps {
		for j, s := range seeds {
			seeds[j] = getMatch(s, m)
		}
		fmt.Println(seeds)
	}
	sort.Ints(seeds)
	fmt.Println(seeds[0])
}

func getMatch(input int, m []mapElement) int {
	for _, v := range m {
		sourceMax := v.Source + v.Range - 1
		sourceMin := v.Source
		if input <= sourceMax && input >= sourceMin {
			return v.Destination - (v.Source - input)
		}
	}
	return input
}

func part2Step(startRanges [][2]int, m []mapElement) [][2]int {
	var newStartRanges [][2]int
	for len(startRanges) > 0 {
		currentStartRange := startRanges[0]
		newStartRange, remainStartRange := getMatchPart2(currentStartRange, m)
		newStartRanges = append(newStartRanges, newStartRange)
		startRanges = startRanges[1:]
		for _, remain := range remainStartRange {
			if remain[0] > -1 {
				startRanges = append(startRanges, remain)
			}
		}
	}
	//printSeeds(newStartRanges)
	return newStartRanges
}

func printSeeds(newStartRanges [][2]int) []int {
	var seeds []int
	for _, v := range newStartRanges {
		for i := 0; i < v[1]; i++ {
			seeds = append(seeds, v[0]+i)
		}
	}
	fmt.Println(seeds)
	return seeds
}

func getMatchPart2(startRange [2]int, m []mapElement) ([2]int, [][2]int) {
	var newStartRange [2]int
	for _, v := range m {
		sourceMax := v.Source + v.Range - 1
		sourceMin := v.Source
		if startRange[0] >= sourceMin && startRange[0] <= sourceMax {
			newStartRange[0] = v.Destination - (v.Source - startRange[0])
			if startRange[0]+startRange[1] <= v.Source+v.Range {
				newStartRange[1] = startRange[1]
				return newStartRange, [][2]int{{-1, -1}}
			} else {
				newStartRange[1] = v.Source + v.Range - startRange[0]
				return newStartRange, [][2]int{{startRange[0] + newStartRange[1], startRange[1] - newStartRange[1]}}
			}

		}
	}
	for _, v := range m {
		if startRange[0] < v.Source && startRange[0]+startRange[1] > v.Source {
			newStartRange[0] = v.Destination
			newStartRange[1] = startRange[1] - v.Source + startRange[0]
			return newStartRange, [][2]int{{startRange[0], startRange[1] - newStartRange[1]}}
		}
		if startRange[0] > v.Source && startRange[0] < v.Source+v.Range {
			newStartRange[0] = v.Destination - (v.Source - startRange[0])
			newStartRange[1] = v.Source + v.Range - startRange[0]
			return newStartRange, [][2]int{{v.Source + v.Range, startRange[1] - newStartRange[1]}}
		}
		if startRange[0] < v.Source && startRange[0]+startRange[1] > v.Source+v.Range {
			newStartRange[0] = v.Destination
			newStartRange[1] = v.Range
			return newStartRange, [][2]int{{startRange[0], v.Source - startRange[0]}, {v.Source + v.Range, startRange[1] - newStartRange[1] - v.Source - startRange[0]}}
		}
	}
	return startRange, [][2]int{{-1, -1}}
}
