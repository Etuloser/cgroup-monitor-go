package main

import (
	"flag"
	"log"
	"time"

	"cgroup-monitor-go/internal/event"
	"cgroup-monitor-go/internal/monitor"
)

func main() {
	path := flag.String("cgroup", "/sys/fs/cgroup", "cgroup v2 path")
	interval := flag.Duration("interval", 1*time.Second, "poll interval")
	flag.Parse()

	ch := make(chan event.Event, 10)
	m := monitor.New(*path, *interval)

	go m.Run(ch)

	log.Println("cgroup monitor started:", *path)
	for e := range ch {
		log.Printf("[%s] %s\n", e.Type, e.Message)
	}
}
