package monitoring

import (
	"time"

	psutilCPU "github.com/shirou/gopsutil/v3/cpu"
	psutilMEM "github.com/shirou/gopsutil/v3/mem"
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
		cpu    bool
		mem    bool
		output chan []Usage
		quit   chan bool
	}
)

const USAGE_SIZE int = 4 + 1

func NewMonitor(msPeriod int, cpu, mem bool) Monitor {
	return &monitor{ticker: time.NewTicker(time.Duration(msPeriod) * time.Millisecond),
		cpu:    cpu,
		mem:    mem,
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
				usages := getUsages(m.cpu, m.mem)
				if len(usages) > 1 {
					m.output <- usages
				}
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

func getUsages(cpu, mem bool) []Usage {
	// TODO add other metrics
	usage := []Usage{}
	if cpu {
		p, _ := psutilCPU.Percent(0, false)
		usage = append(usage, Usage{Name: [4]byte{'c', 'p', 'u', ' '}, Percent: p[0]})
	}
	if mem {
		m, _ := psutilMEM.VirtualMemory()
		usage = append(usage, Usage{Name: [4]byte{'m', 'e', 'm', ' '}, Percent: m.UsedPercent})
	}
	return usage
}
