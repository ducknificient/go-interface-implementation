package main

type Laptop interface {
	PowerOn()
}

func main() {
	//gpu := graphicCard{"RTX 2060", 4096}
	//cpu := Processor{Name: "Intel Core i7 10900K"}
	//fmt.Println(gpu.turnOn())
	//fmt.Println(cpu.turnOn())

	var alienware Laptop
	alienware = Processor{"Intel Core i9 10900K", 8, 16, 4.0}
	alienware = graphicCard{"RTX 2060", 6132}
	alienware.PowerOn()
}
