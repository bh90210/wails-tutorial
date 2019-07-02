package monitoring

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// GetCPU .
func GetCPU() (int32, float64, float64, float64, float64) {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
	}

	times, err := cpu.Times(false)
	if err != nil {
	}

	user := times[0].User
	system := times[0].System
	idle := times[0].Idle
	nice := times[0].Nice

	return int32(math.Round(percent[0])), user, system, idle, nice
}

// GetDISKCounters .
func GetDISKCounters() (float64, float64) {
	iocounters, err := disk.Usage("/")
	if err != nil {
	}

	usedPercent := iocounters.UsedPercent
	inodesUsedPercent := iocounters.InodesUsedPercent

	return usedPercent, inodesUsedPercent
}

// GetLoad .
func GetLoad() (float64, float64, float64, int32, int32, int32) {
	avgstat, err := load.Avg()
	if err != nil {
	}

	load1 := avgstat.Load1
	load5 := avgstat.Load5
	load15 := avgstat.Load15

	miscstat, err := load.Misc()
	if err != nil {
	}

	procsRunning := miscstat.ProcsRunning
	procsBlocked := miscstat.ProcsBlocked
	ctxt := miscstat.Ctxt

	return load1, load5, load15, int32(procsRunning), int32(procsBlocked), int32(ctxt)
}

// GetMem .
func GetMem() (uint64, uint64, uint64, uint64, uint64) {
	swapMemory, err := mem.SwapMemory()
	if err != nil {
	}

	total := swapMemory.Total
	used := swapMemory.Used
	free := swapMemory.Free
	sin := swapMemory.Sin
	sout := swapMemory.Sout

	return total, used, free, sin, sout
}
