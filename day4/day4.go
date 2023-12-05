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
	matches := make([]int, len(lines))
	cardAmounts := make([]int, len(lines))
	for i, l := range lines {
		matches[i] = getMatches(l)
		cardAmounts[i] = 1
	}
	part1(matches)
	part2(cardAmounts, matches)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part2(cardAmounts []int, matches []int) {
	var sum int
	for i, c := range cardAmounts {
		sum += c
		for j := 1; j <= matches[i]; j++ {
			cardAmounts[i+j] += c
		}
	}
	fmt.Println(sum)
}

func part1(matches []int) {
	var sum int
	for _, m := range matches {
		var cardPoints int
		for i := 0; i < m; i++ {
			if cardPoints == 0 {
				cardPoints = 1
			} else {
				cardPoints *= 2
			}
		}
		sum += cardPoints
	}
	fmt.Println(sum)
}

func getMatches(game string) int {
	var matches int
	leftRight := strings.Split(game[strings.Index(game, ": ")+2:], " | ")
	leftSet := map[string]bool{}
	for i := 0; i < len(leftRight[0]); i += 3 {
		leftSet[leftRight[0][i:i+2]] = true
	}
	for i := 0; i < len(leftRight[1]); i += 3 {
		var n = leftRight[1][i : i+2]
		if leftSet[n] {
			matches++
		}
	}
	return matches
}
