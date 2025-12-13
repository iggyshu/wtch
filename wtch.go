package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
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

	fmt.Printf("Total:      %d MB\n", totalMB)
	fmt.Printf("Available:  %d MB\n", availableMB)
	fmt.Printf("Used:       %.1f %%\n", usedPercent)
}

func watch() {
	for {
		clearScreen()
		fmt.Println("wtch")
		displayMemory()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	watch()
}
