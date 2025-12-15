package internal

import (
	"fmt"
	"log"
	"strconv"

	"github.com/c9s/goprocinfo/linux"
	"github.com/ssimunic/gosensors"
)

type CPUStats struct {
	Temperature string
	CoreLoad    []string
}

func GetWorkload() *linux.Stat {
	statistic, err := linux.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	return statistic
}

func CalcCPUStats(curr, prev *linux.Stat) CPUStats {
	NumCPU := GetInfoCpu().NumCPU()
	cpuCores := make([]string, NumCPU)

	for i := 0; i < NumCPU; i++ {
		cpuCores[i] = calcSingleCoreUsage(curr.CPUStats[i], prev.CPUStats[i], true)
	}

	data := CPUStats{}
	data.CoreLoad = cpuCores
	data.Temperature = getTemperatureProcessor("temp1")

	return data
}

func getTemperatureProcessor(keySensors string) string {
	sensors, err := gosensors.NewFromSystem()

	if err != nil {
		panic(err)
	}

	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
			if key == keySensors {
				return value
			}
		}
	}

	return ""
}

func calcSingleCoreUsage(curr, prev linux.CPUStat, isPercentage bool) string {

	PrevIdle := prev.Idle + prev.IOWait
	Idle := curr.Idle + curr.IOWait

	PrevNonIdle := prev.User + prev.Nice + prev.System + prev.IRQ + prev.SoftIRQ + prev.Steal
	NonIdle := curr.User + curr.Nice + curr.System + curr.IRQ + curr.SoftIRQ + curr.Steal

	PrevTotal := PrevIdle + PrevNonIdle
	Total := Idle + NonIdle

	totald := Total - PrevTotal
	idled := Idle - PrevIdle

	CPULoad := (float64(totald) - float64(idled)) / float64(totald)

	if !isPercentage {
		return fmt.Sprintf("%s", strconv.FormatFloat(CPULoad, 'f', 2, 64))
	}

	return fmt.Sprintf("%s%%", strconv.FormatFloat(CPULoad*100, 'f', 2, 64))
}
