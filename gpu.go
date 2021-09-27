package main

import "fmt"

type graphicCard struct {
	Name string
	VRAM int
}

func (gpu graphicCard) turnOn() string {
	status := gpu.Name + ` is ready`
	return status
}

func (gpu graphicCard) PowerOn() {
	fmt.Printf(gpu.Name+" is powering on with %v  VRAM", gpu.VRAM)
}
