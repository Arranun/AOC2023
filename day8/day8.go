package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Node struct {
	name  string
	left  *Node
	right *Node
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	existingNodes := parseTree(lines[2:])
	part1(existingNodes["AAA"], lines)
	//part2(existingNodes, lines)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func part2(existingNodes map[string]*Node, lines []string) {
	activeNodes := []*Node{}
	for k, v := range existingNodes {
		if k[2] == 'A' {
			activeNodes = append(activeNodes, v)
		}
	}
	steps := 0
	directionsIndex := 0
	for checkAllNodesOnZ(activeNodes) {
		for i, _ := range activeNodes {
			activeNodes[i] = step(activeNodes[i], lines[0][directionsIndex])
		}
		if directionsIndex == len(lines[0])-1 {
			directionsIndex = 0
		} else {
			directionsIndex++
		}
		steps++
	}
	fmt.Println(steps)
}

func checkAllNodesOnZ(nodes []*Node) bool {
	for _, n := range nodes {
		if n.name[2] != 'Z' {
			return false
		}
	}
	return true
}

func step(currentNode *Node, direction uint8) *Node {
	if direction == 76 {
		return currentNode.left
	} else {
		return currentNode.right
	}
}

func part1(currentNode *Node, lines []string) {
	steps := 0
	directionsIndex := 0
	for currentNode.name != "ZZZ" {
		currentNode = step(currentNode, lines[0][directionsIndex])

		if directionsIndex == len(lines[0])-1 {
			directionsIndex = 0
		} else {
			directionsIndex++
		}
		steps++
	}
	fmt.Println(steps)
}

func parseTree(lines []string) map[string]*Node {
	nodeMap := map[string][2]string{}
	for _, v := range lines {
		indexRight := strings.Index(v, "(") + 1
		indexLeft := strings.Index(v, ", ") + 2
		nodeMap[v[0:3]] = [2]string{v[indexRight : indexRight+3], v[indexLeft : indexLeft+3]}
	}
	root := Node{name: "AAA"}
	activeNodes := []*Node{&root}
	existingNodes := map[string]*Node{"AAA": &root}
	visitedNodes := map[string]bool{"AAA": true}
	for len(activeNodes) > 0 {
		activeNode := activeNodes[0]
		leftNode := createNode(nodeMap[activeNode.name][0], &existingNodes)
		rightNode := createNode(nodeMap[activeNode.name][1], &existingNodes)
		activeNode.left = leftNode
		activeNode.right = rightNode

		for _, n := range []*Node{activeNode.left, activeNode.right} {
			if n != nil && !visitedNodes[n.name] {
				activeNodes = append(activeNodes, n)
				visitedNodes[n.name] = true
			}
		}
		activeNodes = activeNodes[1:]
		delete(nodeMap, activeNode.name)
	}
	return existingNodes
}

func createNode(name string, existingNodes *map[string]*Node) *Node {
	if (*existingNodes)[name] == nil {
		newNode := &Node{name: name}
		(*existingNodes)[name] = newNode
		return newNode
	} else {
		return (*existingNodes)[name]
	}
}
