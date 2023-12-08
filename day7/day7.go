package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type hand struct {
	cards string
	bid   int
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	hands := make([]hand, len(lines))
	for i, l := range lines {
		split := strings.Split(l, " ")
		hands[i] = hand{split[0], helper.RemoveError(strconv.Atoi(split[1]))}
	}
	sort.Slice(hands[:], func(i, j int) bool {
		leftHandScore := getHandScore(hands[i])
		rightHandScore := getHandScore(hands[j])
		if leftHandScore == rightHandScore {
			index := 0
			for cardToNumber(rune(hands[i].cards[index])) == cardToNumber(rune(hands[j].cards[index])) {
				index++
			}
			cardNumberLeft := cardToNumber(rune(hands[i].cards[index]))
			cardNumberRight := cardToNumber(rune(hands[j].cards[index]))
			return cardNumberLeft < cardNumberRight
		}
		return leftHandScore < rightHandScore
	})
	sum := 0
	for i, v := range hands {
		s := strings.Split(v.cards, "")
		sort.Strings(s)
		fmt.Printf("%s : %s : %d : %d \n", v.cards, strings.Join(s, ""), i+1, getHandScore(v))
		sum += v.bid * (i + 1)
	}
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getHandScore(h hand) int {
	cardMap := map[string]int{}
	highest := 0
	for _, v := range strings.Split(h.cards, "") {
		cardMap[v]++
		if cardMap[v] > highest {
			highest = cardMap[v]
		}
	}
	if highest < 2 {
		return highest
	}
	if highest == 2 {
		pairAmount := 0
		for _, v := range cardMap {
			if v == 2 {
				pairAmount++
			}
		}
		if pairAmount == 1 {
			return 2
		}
		return 3
	}
	if highest > 3 {
		return highest + 2
	}
	for _, v := range cardMap {
		if v == 2 {
			return 5
		}
	}
	return 4
}

func cardToNumber(c rune) int {
	cardToNumberMap := map[rune]int{'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	if cardToNumberMap[c] != 0 {
		return cardToNumberMap[c]
	}
	return int(c - 48)
}
