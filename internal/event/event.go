package event

import "time"

type Type string

const (
	EventOOM         Type = "OOM"
	EventOOMKill     Type = "OOM_Kill"
	EventCPUThrottle Type = "CPU_THROTTLE"
)

type Event struct {
	Type      Type
	Message   string
	Timestamp time.Time
}
