package main

import (
	"AOC2023/helper"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Module interface {
	Interpret(signal bool, sender string, signals *[2]int) []Order
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

type Output struct {
	BaseModule
	signals [2]int
}

type Order struct {
	receiver *Module
	sender   string
	signal   bool
}

func (b *BaseModule) AddOutput(output *Module) {
	b.Output = append(b.Output, output)
}

func (b *Broadcast) AddOutput(output *Module) {
	b.BaseModule.AddOutput(output)
}

func (f *FlipFlop) AddOutput(output *Module) {
	f.BaseModule.AddOutput(output)
}

func (c *Conjunction) AddOutput(output *Module) {
	c.BaseModule.AddOutput(output)
}

func (o *Output) AddOutput(output *Module) {

}

func (b *Broadcast) Interpret(input bool, sender string, signals *[2]int) []Order {
	//fmt.Printf("Sender: %s, Pulse: %b, Receiver: %s\n", sender, input, b.BaseModule.Name)
	if input {
		signals[1]++
	} else {
		signals[0]++
	}
	var orders []Order
	for _, o := range b.Output {
		orders = append(orders, Order{o, b.Name, input})
		//(*o).Interpret(input, b.Name, signals)
	}
	return orders
}

func (f *FlipFlop) Interpret(input bool, sender string, signals *[2]int) []Order {
	//fmt.Printf("Sender: %s, Pulse: %b, Receiver: %s\n", sender, input, f.BaseModule.Name)
	if input {
		signals[1]++
	} else {
		signals[0]++
	}
	var orders []Order
	if !input {
		f.on = !f.on
		for _, o := range f.Output {
			orders = append(orders, Order{o, f.Name, f.on})
			//(*o).Interpret(f.on, f.Name, signals)
		}
	}
	return orders
}

func (c *Conjunction) Interpret(input bool, sender string, signals *[2]int) []Order {
	//fmt.Printf("Sender: %s, Pulse: %b, Receiver: %s\n", sender, input, c.BaseModule.Name)
	if input {
		signals[1]++
	} else {
		signals[0]++
	}
	c.Lastinput[sender] = input
	signal := c.getOutput()
	var orders []Order
	for _, o := range c.Output {
		orders = append(orders, Order{o, c.Name, signal})
		//(*o).Interpret(signal, c.Name, signals)
	}
	return orders
}

func (c *Conjunction) getOutput() bool {
	signal := false
	for _, v := range c.Lastinput {
		if !v {
			signal = true
		}
	}
	return signal
}

func (o *Output) Interpret(input bool, sender string, signals *[2]int) []Order {
	//fmt.Printf("Sender: %s, Pulse: %b, Receiver: %s\n", sender, input, o.BaseModule.Name)
	if input {
		signals[1]++
		o.signals[1]++
	} else {
		signals[0]++
		o.signals[0]++
	}
	return []Order{}
}

func main() {
	args := os.Args[1:]
	lines := helper.ReadTextFile(args[0])
	start := time.Now()
	elapsed := time.Since(start)
	//part1(getModules(lines))
	part2(getModules(lines))
	log.Printf("Took %s", elapsed)
}

func getModules(lines []string) map[string]Module {
	modules := map[string]Module{}
	conjunctions := map[string]bool{}
	for _, l := range lines {
		split := strings.Split(l, " -> ")
		switch split[0][0] {
		case 'b':
			modules["roadcaster"] = &Broadcast{BaseModule: BaseModule{Name: "roadcaster", Output: make([]*Module, 0)}}
		case '%':
			modules[split[0][1:]] = &FlipFlop{BaseModule: BaseModule{Name: split[0][1:], Output: make([]*Module, 0)}}
		case '&':
			modules[split[0][1:]] = &Conjunction{BaseModule: BaseModule{Name: split[0][1:], Output: make([]*Module, 0)}, Lastinput: map[string]bool{}}
			conjunctions[split[0][1:]] = true
		case '!':
			modules[split[0][1:]] = &Output{BaseModule: BaseModule{Name: split[0][1:], Output: make([]*Module, 0)}}
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
	return modules
}

func part2(modules map[string]Module) {
	//signals := [2]int{}
	//var i int
	//for {
	//	var orders []Order
	//	orders = modules["roadcaster"].Interpret(false, "button", &signals)
	//	for len(orders) > 0 {
	//		order := orders[0]
	//		orders = orders[1:]
	//		if modules["zg"].(*Conjunction).Lastinput["lm"] {
	//			break
	//		}
	//		newOrders := (*order.receiver).Interpret(order.signal, order.sender, &signals)
	//		orders = append(orders, newOrders...)
	//	}
	//	i++
	//}
	fmt.Println(helper.LCMArray([]int{3911, 4057, 3907, 3929}))
}

func part1(modules map[string]Module) {
	signals := [2]int{}
	for i := 0; i < 1000; i++ {
		buttonPress(modules, signals)
	}
	fmt.Println(signals[0] * signals[1])
}

func buttonPress(modules map[string]Module, signals [2]int) {
	var orders []Order
	orders = modules["roadcaster"].Interpret(false, "button", &signals)
	for len(orders) > 0 {
		order := orders[0]
		orders = orders[1:]
		newOrders := (*order.receiver).Interpret(order.signal, order.sender, &signals)
		orders = append(orders, newOrders...)
	}
	fmt.Println()
}
