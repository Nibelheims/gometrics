package monitoring

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type (
	Usage struct {
		CpuPercent float64
		MemPercent float64
		// todo network stats
	}

	Monitor interface {
		C() <-chan Usage
		Run()
		Stop()
	}

	monitor struct {
		ticker *time.Ticker
		output chan Usage
		quit   chan bool
	}
)

func NewMonitor(msPeriod int) Monitor {
	return &monitor{ticker: time.NewTicker(time.Duration(msPeriod) * time.Millisecond),
		output: make(chan Usage),
		quit:   make(chan bool)}
}

func (m *monitor) C() <-chan Usage {
	return m.output
}

func (m *monitor) Run() {
	go func() {
		defer m.ticker.Stop()
		for {
			select {
			case <-m.ticker.C:
				usage := getUsage()
				m.output <- usage
			case <-m.quit:
				return
			}
		}
	}()
}

func (m *monitor) Stop() {
	m.quit <- true
	close(m.output)
}

func getUsage() Usage {
	p, _ := cpu.Percent(0, false)
	m, _ := mem.VirtualMemory()
	return Usage{CpuPercent: p[0], MemPercent: m.UsedPercent}
}
