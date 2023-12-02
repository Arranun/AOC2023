package main

import (
	"AOC2023/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var linesOnlyNumbers []string
	for _, l := range lines {
		var numbers string
		for _, c := range l {
			if c > 48 && c < 59 {
				numbers = numbers + string(c)
			}
		}
		linesOnlyNumbers = append(linesOnlyNumbers, numbers)
	}
	var sum int
	for _, l := range linesOnlyNumbers {
		sum += helper.RemoveError(strconv.Atoi(l[:1] + l[len(l)-1:]))
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	var sum int
	for _, l := range lines {
		var firstDigit = ""
		var lastDigit = ""
		var index = 0
		for firstDigit == "" {
			firstDigit = searchDigits(l[:index])
			index++
		}
		index = 0
		for lastDigit == "" {
			lastDigit = searchDigits(l[len(l)-index:])
			index++
		}
		sum += helper.RemoveError(strconv.Atoi(firstDigit + lastDigit))
	}
	fmt.Println(sum)
}

func searchDigits(input string) string {
	var digits = []string{
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	}
	for i, d := range digits {
		if strings.Contains(input, d) {
			if i%2 == 0 {
				return digits[i+1]
			}
			return digits[i]

		}
	}
	return ""
}
