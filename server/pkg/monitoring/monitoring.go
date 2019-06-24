package monitoring

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
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
