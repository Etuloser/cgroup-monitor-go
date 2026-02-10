package cgroup

import (
	"strconv"
	"strings"
)

type MemoryEvents struct {
	OOM     int64
	OOMKill int64
	High    int64
}

func ReadMemoryEvents(path string) (MemoryEvents, error) {
	data, err := ReadFile(path)
	if err != nil {
		return MemoryEvents{}, err
	}
	var ev MemoryEvents
	for _, line := range strings.Split(data, "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		switch fields[0] {
		case "oom":
			ev.OOM, _ = parse(fields[1])
		case "oom_kill":
			ev.OOMKill, _ = parse(fields[1])
		case "high":
			ev.High, _ = parse(fields[1])
		}
	}
	return ev, nil
}

func parse(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
