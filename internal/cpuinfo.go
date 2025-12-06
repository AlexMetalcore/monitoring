package internal

import (
	"log"

	"github.com/c9s/goprocinfo/linux"
)

func GetInfoCpu() *linux.CPUInfo {
	info, err := linux.ReadCPUInfo("/proc/cpuinfo")

	if err != nil {
		log.Fatal("cpuinfo read fail")
	}

	return info
}
