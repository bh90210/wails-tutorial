package monitoring

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	_ "github.com/shirou/gopsutil/disk"
	_ "github.com/shirou/gopsutil/load"
	_ "github.com/shirou/gopsutil/mem"
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
