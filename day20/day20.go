package main

import (
	"AOC2023/helper"
	"log"
	"os"
	"strings"
	"time"
)

type Module interface {
	Interpret(signal bool, sender string)
	AddOutput(module *Module)
}

type BaseModule struct {
	Name   string
	Output []*Module
}

type FlipFlop struct {
	BaseModule
	on bool
}

type Conjunction struct {
	BaseModule
	Lastinput map[string]bool
}

type Broadcast struct {
	BaseModule
}

func (b BaseModule) AddOutput(output *Module) {
	b.Output = append(b.Output, output)
}

func (b Broadcast) AddOutput(output *Module) {
	b.BaseModule.AddOutput(output)
}

func (f FlipFlop) AddOutput(output *Module) {
	f.BaseModule.AddOutput(output)
}

func (c Conjunction) AddOutput(output *Module) {
	c.BaseModule.AddOutput(output)
}

func (b Broadcast) Interpret(input bool, sender string) {
	for _, o := range b.Output {
		(*o).Interpret(input, b.Name)
	}
}

func (f FlipFlop) Interpret(input bool, sender string) {
	if !input {
		f.on = !f.on
		for _, o := range f.Output {
			(*o).Interpret(f.on, f.Name)
		}
	}
}

func (c Conjunction) Interpret(input bool, sender string) {
	c.Lastinput[sender] = input
	signal := false
	for _, v := range c.Lastinput {
		if !v {
			signal = true
		}
	}
	for _, o := range c.Output {
		(*o).Interpret(signal, c.Name)
	}
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	modules := map[string]Module{}
	conjunctions := map[string]bool{}
	for _, l := range lines {
		split := strings.Split(l, " -> ")
		switch split[0][0] {
		case 'b':
			modules["roadcast"] = Broadcast{BaseModule: BaseModule{Name: "roadcast"}}
		case '%':
			modules[split[0][1:]] = FlipFlop{BaseModule: BaseModule{Name: split[0][1:]}}
		case '&':
			modules[split[0][1:]] = Conjunction{BaseModule: BaseModule{Name: split[0][1:]}}
			conjunctions[split[0][1:]] = true
		}

	}
	for _, l := range lines {
		split := strings.Split(l, " -> ")
		outputs := strings.Split(split[1], ", ")
		currentModule := modules[split[0][1:]]
		for _, o := range outputs {
			module := modules[o]
			currentModule.AddOutput(&module)
			if conjunctions[o] {
				modules[o].(*Conjunction).Lastinput[split[0][1:]] = false
			}
		}
	}
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
