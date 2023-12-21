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

type Part struct {
	x, m, a, s      int
	currentWorkFlow *WorkFlow
}

type WorkFlow struct {
	name  string
	rules []*Rule
}

type Rule struct {
	condition   string
	element     string
	number      int
	destination *WorkFlow
}

type Path struct {
	rules           []*Rule
	currentWorkflow *WorkFlow
	partRange       map[string][2]int
}

type Limit struct {
	minMaX map[string][2]int
}

type Node struct {
	rule Rule
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	parts := parse(lines)
	part1(parts)
	partRange := map[string][2]int{}
	partRange["x"] = [2]int{1, 4000}
	partRange["m"] = [2]int{1, 4000}
	partRange["a"] = [2]int{1, 4000}
	partRange["s"] = [2]int{1, 4000}
	activePaths := []Path{{[]*Rule{}, parts[0].currentWorkFlow, partRange}}
	acceptedPaths := []Path{}
	for len(activePaths) > 0 {
		activePath := activePaths[0]
		activePaths = activePaths[1:]
		for _, r := range activePath.currentWorkflow.rules {
			if r.destination.name != "A" {
				newParts, remainingParts := reduceParts(activePath.partRange, r)
				activePath.partRange = remainingParts
				newPath := Path{append(activePath.rules, r), r.destination, newParts}
				if r.destination.name != "R" {
					activePaths = append(activePaths, newPath)
				}
			} else if r.destination.name == "A" {
				newParts, remainingParts := reduceParts(activePath.partRange, r)
				activePath.partRange = remainingParts
				newPath := Path{append(activePath.rules, r), r.destination, newParts}
				acceptedPaths = append(acceptedPaths, newPath)
			}
		}
	}
	sum := 0
	for _, path := range acceptedPaths {
		sum += (path.partRange["x"][1] - path.partRange["x"][0] + 1) * (path.partRange["m"][1] - path.partRange["m"][0] + 1) * (path.partRange["s"][1] - path.partRange["s"][0] + 1) * (path.partRange["a"][1] - path.partRange["a"][0] + 1)
	}
	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func reduceParts(partRange map[string][2]int, r *Rule) (map[string][2]int, map[string][2]int) {
	newPartRange := map[string][2]int{}
	remainingPartRange := map[string][2]int{}
	for k, v := range partRange {
		newPartRange[k] = v
		remainingPartRange[k] = v
	}
	if r.condition == "<" {
		newPartRange[r.element] = [2]int{newPartRange[r.element][0], r.number - 1}
		remainingPartRange[r.element] = [2]int{r.number, remainingPartRange[r.element][1]}
	}
	if r.condition == ">" {
		newPartRange[r.element] = [2]int{r.number + 1, newPartRange[r.element][1]}
		remainingPartRange[r.element] = [2]int{remainingPartRange[r.element][0], r.number}
	}
	return newPartRange, remainingPartRange
}

func findOverlap(a, b [2]int) [2]int {
	maxStart := slices.Max([]int{a[0], b[0]})
	minEnd := slices.Min([]int{a[1], b[1]})
	if maxStart <= minEnd {
		return [2]int{maxStart, minEnd}
	}
	return [2]int{}
}

func (p *Path) getLimit() Limit {
	limit := Limit{map[string][2]int{}}
	for _, el := range []string{"x", "m", "a", "s"} {
		limit.minMaX[el] = [2]int{1, 4000}
	}
	for _, r := range p.rules {
		if r.condition == "<" && (r.number < limit.minMaX[r.element][1] || limit.minMaX[r.element][1] == 0) {
			limit.minMaX[r.element] = [2]int{limit.minMaX[r.element][0], r.number - 1}
		}
		if r.condition == ">" && (r.number > limit.minMaX[r.element][0] || limit.minMaX[r.element][0] == 0) {
			limit.minMaX[r.element] = [2]int{r.number + 1, limit.minMaX[r.element][1]}
		}
	}
	return limit
}

func getPossibleValues(acceptedPaths []Path) [][4]map[int]bool {
	possibleValues := make([][4]map[int]bool, len(acceptedPaths))
	for i, _ := range possibleValues {
		possibleValues[i] = [4]map[int]bool{{}, {}, {}, {}}
	}
	for i, p := range acceptedPaths {
		currentPossibleValues := possibleValues[i]
		for i := 0; i < 4; i++ {
			for j := 1; j < 4001; j++ {
				currentPossibleValues[i][j] = true
			}
		}
		elementToIndexMap := map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}
		for _, r := range p.rules {
			if r.condition == "<" {
				for i := r.number; i < 4001; i++ {
					delete(currentPossibleValues[elementToIndexMap[r.element]], i)
				}
			}
			if r.condition == ">" {
				for i := 0; i < r.number; i++ {
					delete(currentPossibleValues[elementToIndexMap[r.element]], i)
				}
			}
		}
	}
	return possibleValues
}

func combineLimits(limits [][2]int) int {
	var sum int
	for i := 1; i < 4001; i++ {
		for _, limit := range limits {
			if i >= limit[0] && i <= limit[1] {
				sum++
				break
			}
		}
	}
	return sum
}

func getAlwaysMinAndMax(alwaysPossibleMax, alwaysPossibleMin *[4]int, pv *[4]map[int]bool, value int) [2]int {
	for j := 4000; j > 0; j-- {
		if pv[value][j] && alwaysPossibleMax[value] < j {
			alwaysPossibleMax[value] = j
			break
		}
	}
	for j := 1; j < 4001; j++ {
		if pv[value][j] && alwaysPossibleMin[value] > j {
			alwaysPossibleMin[value] = j
			break
		}
	}
	return [2]int{alwaysPossibleMin[value], alwaysPossibleMax[value]}
}

func check(combination [4]int, possibleValues *[][4]map[int]bool) bool {
	for _, pv := range *possibleValues {
		if pv[0][combination[0]] && pv[1][combination[1]] && pv[2][combination[2]] && pv[3][combination[3]] {
			return true
		}
	}
	return false
}

func part1(parts []Part) {
	sum := 0
	for _, p := range parts {
		if p.check() {
			sum += p.x + p.m + p.a + p.s
		}
	}
	fmt.Println(sum)
}

func parse(lines []string) []Part {
	workFlows := map[string]*WorkFlow{}
	lI := 0
	for lines[lI] != "" {
		var w *WorkFlow
		split := strings.Split(lines[lI], "{")
		w = getWorkFlow(&workFlows, split[0])
		ruleStrings := strings.Split(split[1], ",")
		for i := 0; i < len(ruleStrings)-1; i++ {
			rs := ruleStrings[i]
			rss := strings.Split(rs, ":")
			r := Rule{}
			r.element = rss[0][:1]
			r.condition = rss[0][1:2]
			r.number = helper.RemoveError(strconv.Atoi(rss[0][2:len(rss[0])]))
			r.destination = getWorkFlow(&workFlows, rss[1])
			w.rules = append(w.rules, &r)
		}
		r := Rule{}
		rs := ruleStrings[len(ruleStrings)-1][:len(ruleStrings[len(ruleStrings)-1])-1]
		r.destination = getWorkFlow(&workFlows, rs)
		w.rules = append(w.rules, &r)
		lI++
	}
	fmt.Println()
	lI++
	var parts []Part
	for lI < len(lines) {
		split := strings.Split(lines[lI][1:len(lines[lI])-1], ",")
		p := Part{}
		p.x = helper.RemoveError(strconv.Atoi(split[0][2:]))
		p.m = helper.RemoveError(strconv.Atoi(split[1][2:]))
		p.a = helper.RemoveError(strconv.Atoi(split[2][2:]))
		p.s = helper.RemoveError(strconv.Atoi(split[3][2:]))
		p.currentWorkFlow = workFlows["in"]
		parts = append(parts, p)
		lI++
	}
	return parts
}

func (p *Part) check() bool {
	for {
		p.step()
		if p.currentWorkFlow.name == "A" {
			return true
		}
		if p.currentWorkFlow.name == "R" {
			return false
		}
	}
}

func (p *Part) step() {
	for _, rule := range p.currentWorkFlow.rules {
		if p.checkRule(rule) {
			p.currentWorkFlow = rule.destination
			return
		}
	}
}

func (p *Part) checkRule(rule *Rule) bool {
	if rule.condition == "" {
		return true
	}
	partValue := p.getValueByName(rule.element)
	switch rule.condition {
	case "<":
		return partValue < rule.number
	case ">":
		return partValue > rule.number
	}
	return false
}

func (p *Part) getValueByName(name string) int {
	switch name {
	case "x":
		return p.x
	case "m":
		return p.m
	case "s":
		return p.s
	case "a":
		return p.a
	default:
		return -1
	}
}

// Returns true if Rejected or Accepted
func (p *Part) setNewWorkFlow(rule *Rule) bool {
	p.currentWorkFlow = rule.destination
	if p.currentWorkFlow.name == "Accepted" || p.currentWorkFlow.name == "Rejected" {
		return true
	}
	return false
}

func getWorkFlow(workFlows *map[string]*WorkFlow, name string) *WorkFlow {
	v, ok := (*workFlows)[name]
	if ok {
		return v
	} else {
		(*workFlows)[name] = &WorkFlow{name: name}
		return (*workFlows)[name]
	}
}
