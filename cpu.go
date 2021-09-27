package main

import "fmt"

type Processor struct {
	Name       string
	Core       int
	Thread     int
	ClockSpeed float64
}

func (cpu Processor) turnOn() string {
	status := cpu.Name + ` is turning on`
	return status
}

func (cpu Processor) PowerOn() {
	fmt.Println(cpu.Name + " is powering on")
}
