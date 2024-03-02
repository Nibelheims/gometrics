package monitoring

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type (
	Usage struct {
		Name    [4]byte // up to 4 ASCII chars only, no final null byte
		Percent float64
	}

	Monitor interface {
		C() <-chan []Usage
		Run()
		Stop()
	}

	monitor struct {
		ticker *time.Ticker
		output chan []Usage
		quit   chan bool
	}
)

const USAGE_SIZE int = 4 + 1

func NewMonitor(msPeriod int) Monitor {
	return &monitor{ticker: time.NewTicker(time.Duration(msPeriod) * time.Millisecond),
		output: make(chan []Usage),
		quit:   make(chan bool)}
}

func (m *monitor) C() <-chan []Usage {
	return m.output
}

func (m *monitor) Run() {
	go func() {
		defer m.ticker.Stop()
		for {
			select {
			case <-m.ticker.C:
				usages := getUsages()
				m.output <- usages
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

func getUsages() []Usage {
	// TODO add other metrics
	p, _ := cpu.Percent(0, false)
	m, _ := mem.VirtualMemory()
	return []Usage{
		{Name: [4]byte{'c', 'p', 'u', ' '}, Percent: p[0]},
		{Name: [4]byte{'m', 'e', 'm', ' '}, Percent: m.UsedPercent},
	}
}
