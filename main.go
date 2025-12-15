package main

import (
	"fmt"
	cpustatistics "monitoring/internal"
	diskinfo "monitoring/internal"
	memoryinfo "monitoring/internal"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	cpuStatistic := cpustatistics.GetWorkload()
	previousCPUStats := cpustatistics.GetWorkload()

	coreStats := cpustatistics.CalcCPUStats(cpuStatistic, previousCPUStats)

	countCores := len(coreStats.CoreLoad)

	table.SetCell(
		0,
		0,
		tview.NewTableCell("Temperature CPU").
			SetSelectable(false).
			SetTextColor(tcell.ColorRed),
	)

	countSpaceCores := 1

	table.SetCell(
		countSpaceCores,
		0,
		tview.NewTableCell("").
			SetSelectable(false),
	)

	countCores = countCores + countSpaceCores

	for i := 0; i < countCores; i++ {
		i := i + 1
		table.SetCell(
			i+1,
			0,
			tview.NewTableCell("CPU "+strconv.Itoa(i)+" Core Load").
				SetTextColor(tcell.ColorGreen).
				SetSelectable(true),
		)
	}

	countSpaceTotalMemory := countCores + 1
	countTotalMemory := countSpaceTotalMemory + 1
	countAvailableMemory := countTotalMemory + 1
	countActiveMemory := countAvailableMemory + 1

	table.SetCell(
		countSpaceTotalMemory,
		0,
		tview.NewTableCell("").
			SetSelectable(false),
	)

	table.SetCell(
		countTotalMemory,
		0,
		tview.NewTableCell("Total Memory").
			SetTextColor(tcell.ColorYellow).
			SetSelectable(false),
	)
	table.SetCell(
		countAvailableMemory,
		0,
		tview.NewTableCell("Available Memory").
			SetTextColor(tcell.ColorYellow).
			SetSelectable(false),
	)
	table.SetCell(
		countActiveMemory,
		0,
		tview.NewTableCell("Active Memory").
			SetTextColor(tcell.ColorYellow).
			SetSelectable(false),
	)

	countSpaceTotalSpace := countActiveMemory + 1
	countTotalSpace := countSpaceTotalSpace + 1
	countUsedSpace := countTotalSpace + 1
	countFreeSpace := countUsedSpace + 1

	table.SetCell(
		countSpaceTotalSpace,
		0,
		tview.NewTableCell("").
			SetSelectable(false),
	)
	table.SetCell(
		countTotalSpace,
		0,
		tview.NewTableCell("Total Disk Space").
			SetTextColor(tcell.ColorDarkMagenta).
			SetSelectable(false),
	)
	table.SetCell(
		countUsedSpace,
		0,
		tview.NewTableCell("Used Disk Space").
			SetTextColor(tcell.ColorDarkMagenta).
			SetSelectable(false),
	)
	table.SetCell(
		countFreeSpace,
		0,
		tview.NewTableCell("Free Disk Space").
			SetTextColor(tcell.ColorDarkMagenta).
			SetSelectable(false),
	)

	countSpaceUptime := countFreeSpace + 1
	countUptime := countSpaceUptime + 1

	table.SetCell(
		countSpaceUptime,
		0,
		tview.NewTableCell("").
			SetSelectable(false),
	)
	table.SetCell(
		countUptime,
		0,
		tview.NewTableCell("Uptime").
			SetTextColor(tcell.ColorDarkCyan).
			SetSelectable(false),
	)

	go func() {
		for {
			app.QueueUpdateDraw(func() {
				currentCPUStats := cpustatistics.GetWorkload()
				coreStats := cpustatistics.CalcCPUStats(currentCPUStats, previousCPUStats)
				table.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%s", coreStats.Temperature)).
					SetTextColor(tcell.ColorRed),
				)
				table.SetCell(1, 1, tview.NewTableCell(""))

				for key, item := range coreStats.CoreLoad {
					key := key + 1
					table.SetCell(key+1, 1, tview.NewTableCell(item).SetTextColor(tcell.ColorGreen))
				}

				table.SetCell(countSpaceTotalMemory, 1, tview.NewTableCell(""))
				table.SetCell(countTotalMemory, 1, tview.NewTableCell(fmt.Sprintf("%sMB", memoryinfo.GetTotalMemory())).
					SetTextColor(tcell.ColorYellow),
				)
				table.SetCell(countAvailableMemory, 1, tview.NewTableCell(fmt.Sprintf("%sMB", memoryinfo.GetAvailableMemory())).
					SetTextColor(tcell.ColorYellow),
				)
				table.SetCell(countActiveMemory, 1, tview.NewTableCell(fmt.Sprintf("%sMB", memoryinfo.GetActiveMemory())).
					SetTextColor(tcell.ColorYellow),
				)

				table.SetCell(countTotalSpace, 1, tview.NewTableCell(fmt.Sprintf("%sG", diskinfo.GetAllDisk())).
					SetTextColor(tcell.ColorDarkMagenta),
				)
				table.SetCell(countUsedSpace, 1, tview.NewTableCell(fmt.Sprintf("%sG", diskinfo.GetUsedDisk())).
					SetTextColor(tcell.ColorDarkMagenta),
				)
				table.SetCell(countFreeSpace, 1, tview.NewTableCell(fmt.Sprintf("%sG", diskinfo.GetFreeDisk())).
					SetTextColor(tcell.ColorDarkMagenta),
				)

				duration := time.Since(cpuStatistic.BootTime)
				days := int(duration.Hours() / 24)
				hours := int(duration.Hours()) % 24
				minutes := int(duration.Minutes()) % 60
				seconds := int(duration.Seconds()) % 60

				table.SetCell(countUptime, 1, tview.NewTableCell(
					fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds),
				).SetTextColor(tcell.ColorDarkCyan),
				)

				previousCPUStats = currentCPUStats
			})

			time.Sleep(time.Second)
		}
	}()

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
