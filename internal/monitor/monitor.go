package monitor

import (
	"log"
	"time"

	"cgroup-monitor-go/internal/cgroup"
	"cgroup-monitor-go/internal/event"
)

type Monitor struct {
	CgroupPath string
	Interval   time.Duration

	lastMem cgroup.MemoryEvents
	lastCPU cgroup.CPUStat
	inited  bool
}

func New(path string, interval time.Duration) *Monitor {
	return &Monitor{CgroupPath: path, Interval: interval}
}

func (m *Monitor) Run(ch chan<- event.Event) {
	t := time.NewTicker(m.Interval)
	defer t.Stop()

	for range t.C {
		if !m.inited {
			m.initBaseline()
			m.inited = true
			continue
		}
		m.checkMemory(ch)
		m.checkCPU(ch)
	}
}
func (m *Monitor) initBaseline() {
	mem, _ := cgroup.ReadMemoryEvents(m.CgroupPath + "/memory.events")
	cpu, _ := cgroup.ReadCPUStat(m.CgroupPath + "/cpu.stat")
	m.lastMem = mem
	m.lastCPU = cpu
}
func (m *Monitor) checkMemory(ch chan<- event.Event) {
	ev, err := cgroup.ReadMemoryEvents(m.CgroupPath + "/memory.events")
	if err != nil {
		log.Println("read memory.events:", err)
		return
	}
	if ev.OOM > m.lastMem.OOM {
		ch <- event.Event{Type: event.EventOOM, Message: "memory OOM", Timestamp: time.Now()}
	}
	if ev.OOMKill > m.lastMem.OOMKill {
		ch <- event.Event{Type: event.EventOOMKill, Message: "memory OOM kill", Timestamp: time.Now()}
	}
	m.lastMem = ev
}

func (m *Monitor) checkCPU(ch chan<- event.Event) {
	st, err := cgroup.ReadCPUStat(m.CgroupPath + "/cpu.stat")
	if err != nil {
		log.Println("read cpu.stat:", err)
		return
	}
	if st.NrThrottled > m.lastCPU.NrThrottled {
		ch <- event.Event{
			Type:      event.EventCPUThrottle,
			Message:   "cpu throttling increased",
			Timestamp: time.Now(),
		}
	}
	m.lastCPU = st
}
