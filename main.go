package main

import (
	"fmt"
	cpustatistics "monitoring/internal"
	diskinfo "monitoring/internal"
	memoryinfo "monitoring/internal"
	"time"
)

func main() {
	cpuStatistic := cpustatistics.GetWorkload()
	previousCPUStats := cpustatistics.GetWorkload()

	for {
		time.Sleep(time.Second + 1)
		currentCPUStats := cpustatistics.GetWorkload()
		coreStats := cpustatistics.CalcCPUStats(currentCPUStats, previousCPUStats)
		fmt.Println("Temperature Processors: ", coreStats.Temperature)

		for key, item := range coreStats.CoreLoad {
			fmt.Printf("CPU %d Load: %s \n", key, item)
		}

		fmt.Println()
		fmt.Printf("Total Memory: %s GB \n", memoryinfo.GetTotalMemory())
		fmt.Printf("Available Memory: %s GB \n", memoryinfo.GetMemAvailableMemory())
		fmt.Printf("Active Memory: %s GB \n", memoryinfo.GetActiveMemory())

		fmt.Println()

		fmt.Printf("Total Space: %s GB\n", diskinfo.GetAllDisk())
		fmt.Printf("Used Space: %s GB\n", diskinfo.GetUsedDisk())
		fmt.Printf("Free Space: %s GB\n", diskinfo.GetFreeDisk())

		fmt.Println()
		duration := time.Since(cpuStatistic.BootTime)
		days := int(duration.Hours() / 24)
		hours := int(duration.Hours()) % 24
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60
		fmt.Printf("Uptime: %d days, %d hours, %d minutes, %d seconds\n", days, hours, minutes, seconds)
		fmt.Println()
	}
}
