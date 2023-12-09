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
	var sumPart1 int
	var sumPart2 int
	for _, l := range lines {
		prev, next := getNextValue(l)
		sumPart1 += next
		sumPart2 += prev
		fmt.Println(getNextValue(l))
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getNextValue(line string) (int, int) {
	history := [][]int{}
	history = append(history, helper.StringSliceToIntSlice(strings.Fields(line)))
	finished := false
	for !finished {
		currentLevel := history[len(history)-1]
		newLevel := make([]int, len(currentLevel)-1)
		for i := 0; i < len(currentLevel)-1; i++ {
			newLevel[i] = currentLevel[i+1] - currentLevel[i]
		}
		for _, v := range newLevel {
			if v != 0 {
				finished = false
				break
			}
			finished = true
		}
		history = append(history, newLevel)
	}
	for i := len(history) - 2; i > 0; i-- {
		history[i-1] = append(history[i-1], history[i-1][len(history[i-1])-1]+history[i][len(history[i])-1])
		history[i-1] = append([]int{history[i-1][0] - history[i][0]}, history[i-1]...)
	}
	return history[0][0], history[0][len(history[0])-1]
}
