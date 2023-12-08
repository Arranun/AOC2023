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
	part2 := args[1] == "true"
	start := time.Now()
	hands := make([]hand, len(lines))
	for i, l := range lines {
		split := strings.Split(l, " ")
		hands[i] = hand{split[0], helper.RemoveError(strconv.Atoi(split[1]))}
	}
	sort.Slice(hands[:], func(i, j int) bool {
		leftHandScore := getHandScore(hands[i], part2)
		rightHandScore := getHandScore(hands[j], part2)
		if leftHandScore == rightHandScore {
			index := 0
			for cardToNumber(rune(hands[i].cards[index]), part2) == cardToNumber(rune(hands[j].cards[index]), part2) {
				index++
			}
			cardNumberLeft := cardToNumber(rune(hands[i].cards[index]), part2)
			cardNumberRight := cardToNumber(rune(hands[j].cards[index]), part2)
			return cardNumberLeft < cardNumberRight
		}
		return leftHandScore < rightHandScore
	})
	sum := 0
	for i, v := range hands {
		s := strings.Split(v.cards, "")
		sort.Strings(s)
		fmt.Printf("%s : %s : %d : %d \n", v.cards, strings.Join(s, ""), i+1, getHandScore(v, part2))
		sum += v.bid * (i + 1)
	}
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getHandScore(h hand, part2 bool) int {
	joker := 0
	s := strings.Split(h.cards, "")
	sort.Strings(s)
	cards := strings.Join(s, "")
	if part2 {
		cards = strings.Replace(cards, "J", "", -1)
		joker = len(s) - len(cards)
	}
	pairMap := getCardDuplicates(cards)
	scoreWithoutJoker := getScorWithoutJoker(pairMap)
	if joker == 5 || joker == 4 {
		return 7
	}
	if joker == 3 {
		switch scoreWithoutJoker {
		case 2:
			return 7
		case 1:
			return 6
		}
	}
	if joker == 2 {
		switch scoreWithoutJoker {
		case 4:
			return 7
		case 3:
			return 6
		case 2:
			return 6
		case 1:
			return 4
		}
	}
	if joker == 1 {
		switch scoreWithoutJoker {
		case 6:
			return 7
		case 5:
			return 6
		case 4:
			return 6
		case 3:
			return 5
		case 2:
			return 4
		case 1:
			return 2

		}
	}
	return scoreWithoutJoker
}

func getScorWithoutJoker(pairMap map[int]int) int {
	if pairMap[5] > 0 {
		return 7
	}
	if pairMap[4] > 0 {
		return 6
	}
	if pairMap[3] > 0 {
		if pairMap[2] > 0 {
			return 5
		}
		return 4
	}
	if pairMap[2] == 2 {
		return 3
	}
	if pairMap[2] == 1 {
		return 2
	}
	return 1
}

func getCardDuplicates(cards string) map[int]int {
	pairMap := map[int]int{}
	cardDuplicates := 1
	for i := 0; i < len(cards)-1; i++ {
		if cards[i] == cards[i+1] {
			cardDuplicates++
		} else {
			pairMap[cardDuplicates]++
			cardDuplicates = 1
		}
	}
	pairMap[cardDuplicates]++
	return pairMap
}

func cardToNumber(c rune, part2 bool) int {
	cardToNumberMap := map[rune]int{'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	if part2 {
		cardToNumberMap['J'] = 1
	}
	if cardToNumberMap[c] != 0 {
		return cardToNumberMap[c]
	}
	return int(c - 48)
}
