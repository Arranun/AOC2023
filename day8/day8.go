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
	part2(existingNodes, lines)
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
	stepsToReachZ := []int{}
	for _, n := range activeNodes {
		stepsToReachZ = append(stepsToReachZ, getStepsToZForStartNode(n, lines))
	}
	fmt.Println(helper.LCMArray(stepsToReachZ))
}

func getStepsToZForStartNode(node *Node, lines []string) int {
	steps := 0
	directionsIndex := 0
	for node.name[2] != 'Z' {
		node = step(node, lines[0][directionsIndex])
		if directionsIndex == len(lines[0])-1 {
			directionsIndex = 0
		} else {
			directionsIndex++
		}
		steps++
	}
	return steps
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
	roots := []string{}
	for _, v := range lines {
		indexRight := strings.Index(v, "(") + 1
		indexLeft := strings.Index(v, ", ") + 2
		nodeMap[v[0:3]] = [2]string{v[indexRight : indexRight+3], v[indexLeft : indexLeft+3]}
		if v[2] == 'A' {
			roots = append(roots, v[0:3])
		}
	}
	existingNodes := map[string]*Node{}
	visitedNodes := &map[string]bool{}
	for _, r := range roots {
		parseTreeFromRoot(existingNodes, nodeMap, visitedNodes, r)
	}
	return existingNodes
}

func parseTreeFromRoot(existingNodes map[string]*Node, nodeMap map[string][2]string, visitedNodes *map[string]bool, rootName string) {
	root := Node{name: rootName}
	existingNodes[root.name] = &root
	activeNodes := []*Node{&root}
	(*visitedNodes)[root.name] = true
	for len(activeNodes) > 0 {
		activeNode := activeNodes[0]
		leftNode := createNode(nodeMap[activeNode.name][0], &existingNodes)
		rightNode := createNode(nodeMap[activeNode.name][1], &existingNodes)
		activeNode.left = leftNode
		activeNode.right = rightNode

		for _, n := range []*Node{activeNode.left, activeNode.right} {
			if n != nil && !(*visitedNodes)[n.name] {
				activeNodes = append(activeNodes, n)
				(*visitedNodes)[n.name] = true
			}
		}
		activeNodes = activeNodes[1:]
		delete(nodeMap, activeNode.name)
	}
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
