package main

import (
  "fmt"
  "flag"
  "runtime"
  "time"

  "github.com/karlpokus/filew"
)

var (
  mem = flag.Bool("mem", false, "toggle mem usage output")
  fpath = flag.String("p", "", "path to dir to watch") // required
)

func main() {
  flag.Parse()
  if *fpath == "" {
    panic("path to dir missing")
  }
  if *mem {
    go memUsage()
  }
  events, err := filew.Watch(*fpath, nil)
  if err != nil {
    panic(err)
  }
  fmt.Printf("watching %s\n", *fpath)
  for ev := range events {
    fmt.Println(ev)
  }
}

func memUsage() {
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
