package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	times := helper.StringSliceToIntSlice(strings.Fields(lines[0][strings.Index(lines[1], ":  ")+3:]))
	distances := helper.StringSliceToIntSlice(strings.Fields(lines[1][strings.Index(lines[1], ":  ")+3:]))
	mult := 1
	for i, t := range times {
		mult *= getPossibleTimeButtonPressedWhereDistanceIsHigherThenRecord(float64(t), float64(distances[i]))
	}
	fmt.Println(mult)
	var timePart2 string
	var distancePart2 string
	for i, t := range times {
		timePart2 += strconv.Itoa(t)
		distancePart2 += strconv.Itoa(distances[i])
	}
	fmt.Println(getPossibleTimeButtonPressedWhereDistanceIsHigherThenRecord(float64(helper.RemoveError(strconv.Atoi(timePart2))), float64(helper.RemoveError(strconv.Atoi(distancePart2)))))

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func getPossibleTimeButtonPressedWhereDistanceIsHigherThenRecord(raceTime float64, record float64) int {
	squareResult := math.Sqrt(math.Pow(raceTime, 2) - 4*(record+1))
	x1 := math.Floor((-raceTime - squareResult) / -2)
	x2 := math.Ceil((-raceTime + squareResult) / -2)
	return int(math.Abs(float64(x1-x2)) + 1)
}
