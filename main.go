package main

import (
  "fmt"
  "os"
  "time"
)

// watch watches a file and returns a channels with change events
// and an err if the file does not exist
func watch(fpath string) (events chan string, err error) {
  ogfi, err := os.Stat(fpath)
  if err != nil {
    return
  }
  events = make(chan string)
  go func(){
    ticker := time.NewTicker(1 * time.Second)
    for _ = range ticker.C {
      fi, err := os.Stat(fpath)
      if err != nil {
        events <- "file removed"
        close(events)
        ticker.Stop()
        return
      }
      if fi.Size() != ogfi.Size() {
        ogfi = fi
        events <- "file changed"
      }
    }
  }()
  return
}

func main() {
  fpath := "testdata/file1"
  events, err := watch(fpath)
  if err != nil {
    panic(err)
  }
  fmt.Printf("watching %s\n", fpath)
  for e := range events {
    fmt.Println(e)
  }
}
