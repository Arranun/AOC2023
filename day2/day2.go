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

type cubeset struct {
	color  string
	amount int
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	var games [][][]cubeset
	games = parseGames(lines, games)
	part1part2(games)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part1part2(games [][][]cubeset) {
	var sumPart1 int
	var sumPart2 int
	for i, g := range games {
		if gamePossible(g) {
			sumPart1 += i + 1
		}
		sumPart2 += getGamePower(g)
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}

func gamePossible(game [][]cubeset) bool {
	for _, d := range game {
		var limiter = map[string]int{"red": 12, "green": 13, "blue": 14}
		for _, c := range d {
			limiter[c.color] -= c.amount
			if limiter[c.color] < 0 {
				return false
			}
		}
	}
	return true
}

func getGamePower(game [][]cubeset) int {
	var minimumDice = map[string]int{}
	for _, d := range game {
		minimumDiceCopy := make(map[string]int)
		for k, v := range minimumDice {
			minimumDiceCopy[k] = v
		}
		for _, c := range d {
			minimumDiceCopy[c.color] -= c.amount
		}
		for k, v := range minimumDiceCopy {
			if v < 0 {
				minimumDice[k] -= v
			}
		}
	}
	var sum = 1
	for _, v := range minimumDice {
		sum *= v
	}
	return sum
}

func parseGames(lines []string, games [][][]cubeset) [][][]cubeset {
	for _, g := range lines {
		var game [][]cubeset
		for _, d := range strings.Split(g[strings.Index(g, ": ")+2:], "; ") {
			var draw []cubeset
			for _, c := range strings.Split(d, ", ") {
				var indexSpace = strings.Index(c, " ")
				draw = append(
					draw, cubeset{color: c[indexSpace+1:], amount: helper.RemoveError(strconv.Atoi(c[:indexSpace]))})
			}
			game = append(game, draw)
		}
		games = append(games, game)
	}
	return games
}
