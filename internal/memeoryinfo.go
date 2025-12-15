package internal

import (
	"log"
	"strconv"

	"github.com/c9s/goprocinfo/linux"
)

func GetMemoryInfo() *linux.MemInfo {
	info, err := linux.ReadMemInfo("/proc/meminfo")

	if err != nil {
		log.Fatal("meminfo read fail")
	}

	return info
}

func GetTotalMemory() string {
	info := GetMemoryInfo()
	return strconv.FormatFloat(float64(info.MemTotal/1024), 'f', 2, 64)
}

func GetAvailableMemory() string {
	info := GetMemoryInfo()
	return strconv.FormatFloat(float64(info.MemAvailable/1024), 'f', 2, 64)
}

func GetActiveMemory() string {
	info := GetMemoryInfo()
	return strconv.FormatFloat(float64(info.Active/1024), 'f', 2, 64)
}
