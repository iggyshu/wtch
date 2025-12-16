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
	usedMessage := fmt.Sprintf("Used:       %.1f %%", usedPercent)
	fmt.Println(percentToColor(usedPercent, usedMessage))

	fmt.Println()
}

func displayCpu() {
	p, _ := cpu.Percent(0, true)
	infos, _ := cpu.Info()
	cpuInfo := infos[0]

	fmt.Println("CPU")
	fmt.Printf("Model: %v, Cores: %v, Mhz: %v\n", cpuInfo.ModelName, cpuInfo.Cores, cpuInfo.Mhz)
	for i, percent := range p {
		info := fmt.Sprintf("Core %v: %.1f %%", i+1, percent)
		fmt.Println(percentToColor(percent, info))
	}
	fmt.Println()
}

func percentToColor(percent float64, message string) string {
	switch {
	case percent >= 75:
		return fmt.Sprintf("\033[0;31m%v\033[0m", message) // red
	case percent >= 35:
		return fmt.Sprintf("\033[0;33m%v\033[0m", message) // yellow
	default:
		return message
	}
}

func watch() {
	for {
		clearScreen()
		displayHeader()
		displayMemory()
		displayCpu()

		time.Sleep(1 * time.Second)
	}
}

func main() {
	watch()
}
