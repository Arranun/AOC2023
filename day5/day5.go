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
	seedsPart1 := seeds
	seedsPart2 := []int{}
	for i := 0; i < len(seeds); i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			seedsPart2 = append(seedsPart2, seeds[i]+j)
		}
	}
	getLowestPossibleSoil(maps, seedsPart1)
	getLowestPossibleSoil(maps, seedsPart2)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getLowestPossibleSoil(maps [][]mapElement, seeds []int) {
	for _, m := range maps {
		for j, s := range seeds {
			seeds[j] = getMatch(s, m)
		}
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
