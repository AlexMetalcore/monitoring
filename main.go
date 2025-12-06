package main

import (
	"fmt"
	cpustatistics "monitoring/internal"
	diskinfo "monitoring/internal"
	memoryinfo "monitoring/internal"
	"time"

	"github.com/fatih/color"
)

func main() {
	cpuStatistic := cpustatistics.GetWorkload()
	previousCPUStats := cpustatistics.GetWorkload()

	for {
		time.Sleep(time.Second + 1)
		currentCPUStats := cpustatistics.GetWorkload()
		coreStats := cpustatistics.CalcCPUStats(currentCPUStats, previousCPUStats)
		color.Red("Temperature Processors: %s", coreStats.Temperature)
		fmt.Println()

		for key, item := range coreStats.CoreLoad {
			color.Green("CPU %d Load: %s \n", key, item)
		}

		fmt.Println()
		color.Yellow("Total Memory: %s GB \n", memoryinfo.GetTotalMemory())
		color.Yellow("Available Memory: %s GB \n", memoryinfo.GetMemAvailableMemory())
		color.Yellow("Active Memory: %s GB \n", memoryinfo.GetActiveMemory())

		fmt.Println()

		color.Magenta("Total Space: %s GB\n", diskinfo.GetAllDisk())
		color.Magenta("Used Space: %s GB\n", diskinfo.GetUsedDisk())
		color.Magenta("Free Space: %s GB\n", diskinfo.GetFreeDisk())

		fmt.Println()
		duration := time.Since(cpuStatistic.BootTime)
		days := int(duration.Hours() / 24)
		hours := int(duration.Hours()) % 24
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60

		color.Cyan("Uptime: %d days, %d hours, %d minutes, %d seconds\n", days, hours, minutes, seconds)
		fmt.Println()
	}
}
