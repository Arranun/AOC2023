package main

import (
	"AOC2023/helper"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	partLocations := map[[2]int]string{}
	numbersInInput := map[[2]int]int{}
	partNumbers := map[[2]int]bool{}
	var sum int
	var sumPart2 int
	for x, l := range lines {
		var y = 0
		for y < len(l) {
			var s = l[y]
			if s != '.' && !(s > 47 && s < 58) {
				partLocations[[2]int{x, y}] = string(s)
				y++
			} else if s > 47 && s < 58 {
				var numberEnd = y
				for numberEnd+1 < len(l) && l[numberEnd+1] > 47 && l[numberEnd+1] < 58 {
					numberEnd++
				}
				numbersInInput[[2]int{x, y}] = helper.RemoveError(strconv.Atoi(l[y : numberEnd+1]))
				y = numberEnd + 1
			} else if s == '.' {
				y++
			}
		}
	}
	for p, v := range partLocations {
		directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
		hit := map[[2]int]bool{}
		for _, d := range directions {
			var x = p[0] + d[0]
			var y = p[1] + d[1]
			if x > -1 && y > -1 && len(lines) > x && len(lines[x]) > y && lines[x][y] > 47 && lines[x][y] < 58 {
				hit[[2]int{x, y}] = true
			}
		}
		if v == "*" {
			hitNumbers := map[[2]int]bool{}
			for k, v := range numbersInInput {
				var numberLength = len(strconv.Itoa(v))
				for i := 0; i < numberLength; i++ {
					if hit[[2]int{k[0], k[1] + i}] {
						hitNumbers[k] = true
						break
					}
				}
			}
			if len(hitNumbers) == 2 {
				var mult = 1
				for h, _ := range hitNumbers {
					mult *= numbersInInput[h]
				}
				sumPart2 += mult
			}
		}

		for k, v := range numbersInInput {
			var numberLength = len(strconv.Itoa(v))
			for i := 0; i < numberLength; i++ {
				if hit[[2]int{k[0], k[1] + i}] {
					partNumbers[k] = true
					break
				}
			}
		}
	}
	for k, _ := range partNumbers {
		sum += numbersInInput[k]
	}
	fmt.Println(sum)
	fmt.Println(sumPart2)
}
