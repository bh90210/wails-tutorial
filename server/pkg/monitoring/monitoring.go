package monitoring

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	_ "github.com/shirou/gopsutil/disk"
	_ "github.com/shirou/gopsutil/load"
	_ "github.com/shirou/gopsutil/mem"
)

// GetCPUUsage .
func (s *Stats) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	}

	return &CPUUsage{
		Average: int(math.Round(percent[0])),
	}
}
