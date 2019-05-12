package mem

import (
	"fmt"
	"runtime"
	"time"
)

func Usage() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("memUsage: heap=%vMB, sys=%vMB\n", bToMb(m.Alloc), bToMb(m.Sys))
		time.Sleep(10 * time.Second)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
