package internal

import (
	"log"
	"strconv"

	"github.com/c9s/goprocinfo/linux"
)

func getInfoDisk() *linux.Disk {
	info, err := linux.ReadDisk("/")

	if err != nil {
		log.Fatal("diskinfo read fail")
	}

	return info
}

func GetAllDisk() string {
	info := getInfoDisk()
	return strconv.FormatFloat(float64(info.All/1024/1024/1024), 'f', 2, 64)
}

func GetUsedDisk() string {
	info := getInfoDisk()
	return strconv.FormatFloat(float64(info.Used/1024/1024/1024), 'f', 2, 64)
}

func GetFreeDisk() string {
	info := getInfoDisk()
	return strconv.FormatFloat(float64(info.Free/1024/1024/1024), 'f', 2, 64)
}
