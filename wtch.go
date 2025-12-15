package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func displayHeader() {
	fmt.Println("wtch")
	fmt.Println()
}

func displayMemory() {
	const (
		KB = 1024
		MB = 1024 * KB
	)

	v, _ := mem.VirtualMemory()
	totalMB := v.Total / MB
	availableMB := v.Available / MB
	usedPercent := v.UsedPercent

	fmt.Println("MEMORY")
	fmt.Printf("Total:      %d MB\n", totalMB)
	fmt.Printf("Available:  %d MB\n", availableMB)
	fmt.Printf("Used:       %.1f %%\n", usedPercent)
	fmt.Println()
}

func displayCPU() {
	p, _ := cpu.Percent(0, true)
	infos, _ := cpu.Info()
	cpuInfo := infos[0]

	fmt.Println("CPU")
	fmt.Printf("Model: %v, Cores: %v, Mhz: %v\n", cpuInfo.ModelName, cpuInfo.Cores, cpuInfo.Mhz)
	for i, percent := range p {
		fmt.Printf("Core %v: %.1f %%\n", i+1, percent)
	}
	fmt.Println()
}

func watch() {
	for {
		clearScreen()
		displayHeader()
		displayMemory()
		displayCPU()

		time.Sleep(1 * time.Second)
	}
}

func main() {
	watch()
}
