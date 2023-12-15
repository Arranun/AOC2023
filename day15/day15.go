package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Lens struct {
	label       string
	focalLenght string
}

type Box struct {
	lenses []Lens
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	steps := strings.Split(lines[0], ",")
	part1(steps)
	part2(steps)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part2(steps []string) {
	boxes := make([]Box, 256)
	for _, step := range steps {
		operationIndex := strings.IndexAny(step, "=-")
		label := step[:operationIndex]
		operation := step[operationIndex]
		focalLength := step[operationIndex+1:]
		boxNumber := hash(label)
		if operation == '-' {
			boxes[boxNumber].removeLens(label)
		} else {
			boxes[boxNumber].addLens(Lens{label, focalLength})
		}
	}
	sum := 0
	for i, box := range boxes {
		for j, lens := range box.lenses {
			sum += (i + 1) * (j + 1) * helper.RemoveError(strconv.Atoi(lens.focalLenght))
		}
	}
	fmt.Println(sum)
}

func part1(steps []string) {
	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	fmt.Println(sum)
}

func (b *Box) removeLens(label string) {
	lensIndex := slices.IndexFunc(b.lenses, func(l Lens) bool { return l.label == label })
	if lensIndex != -1 {
		b.lenses = slices.Delete(b.lenses, lensIndex, lensIndex+1)
	}
}

func (b *Box) addLens(lens Lens) {
	lensIndex := slices.IndexFunc(b.lenses, func(l Lens) bool { return l.label == lens.label })
	if lensIndex == -1 {
		b.lenses = append(b.lenses, lens)
	} else {
		b.lenses[lensIndex] = lens
	}
}

func hash(input string) int {
	var cv int32
	for _, r := range input {
		cv += r
		cv *= 17
		cv %= 256
	}
	return int(cv)
}
