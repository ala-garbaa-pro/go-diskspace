package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	// Parse command-line arguments
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: ./diskspace.exe <drive letter>")
		os.Exit(1)
	}

	path := args[0] + ":\\"

	// Get disk usage
	usage, err := disk.Usage(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert bytes to gigabytes (GB)
	totalGB := float64(usage.Total) / (1024 * 1024 * 1024)
	freeGB := float64(usage.Free) / (1024 * 1024 * 1024)
	usedGB := float64(usage.Used) / (1024 * 1024 * 1024)

	// Print disk usage information
	fmt.Printf("Drive %s:\n", args[0])
	fmt.Printf("Total space: %.2fGB\n", totalGB)
	fmt.Printf("Free space: %.2fGB\n", freeGB)
	fmt.Printf("Used space: %.2fGB\n", usedGB)
}
