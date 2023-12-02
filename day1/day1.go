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
	var linesOnlyNumbers []string
	for _, l := range lines {
		var numbers string
		for _, c := range writtenNumberToDigit(l) {
			if c > 48 && c < 59 {
				numbers = numbers + string(c)
			}
		}
		linesOnlyNumbers = append(linesOnlyNumbers, numbers)
	}
	var sum int
	for i, l := range linesOnlyNumbers {
		fmt.Printf("%s : %s \n", lines[i], l)
		sum += helper.RemoveError(strconv.Atoi(l[:1] + l[len(l)-1:]))
	}
	fmt.Println(sum)
}

func writtenNumberToDigit(input string) string {
	r := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9")
	return r.Replace(input)
}
