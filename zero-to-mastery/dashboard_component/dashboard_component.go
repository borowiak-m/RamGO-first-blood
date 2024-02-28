package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (band *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(band.amount); i++ {
		sum += int(band.amount[i])
	}
	return sum / len(band.amount)
}

func (cpu *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for i := 0; i < len(cpu.temp); i++ {
		sum += int(cpu.temp[i])
	}
	return sum / len(cpu.temp)
}

func (mem *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(mem.amount); i++ {
		sum += int(mem.amount[i])
	}
	return sum / len(mem.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{5000, 10000, 13000, 8000, 9000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	mainDash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	fmt.Printf("Average bandwidth usage: %v\n", mainDash.AverageBandwidth())
	fmt.Printf("Average cpu temperature: %v\n", mainDash.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", mainDash.AverageMemoryUsage())
}
