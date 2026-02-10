package cgroup

import (
	"strings"
)

type CPUStat struct {
	NrPeriods   int64
	NrThrottled int64
	ThrottledUS int64
}

func ReadCPUStat(path string) (CPUStat, error) {
	data, err := ReadFile(path)
	if err != nil {
		return CPUStat{}, err
	}
	var st CPUStat
	for _, line := range strings.Split(data, "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		switch fields[0] {
		case "nr_periods":
			st.NrPeriods, _ = parse(fields[1])
		case "nr_throttled":
			st.NrThrottled, _ = parse(fields[1])
		case "throttled_usec":
			st.ThrottledUS, _ = parse(fields[1])
		}
	}
	return st, nil
}
