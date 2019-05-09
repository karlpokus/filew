package filew

import (
  "os"
  "time"
)

type finfo interface {
  Stat(string) (os.FileInfo, error)
}

type ogStat struct {}

func (o ogStat) Stat(fp string) (os.FileInfo, error) {
  return os.Stat(fp)
}

// Watch watches a file and returns a channels with file change events
// and an err if the file does not exist
func Watch(fp string, f finfo) (events chan string, err error) {
  if f == nil {
    f = ogStat{}
  }
  ogfi, err := f.Stat(fp)
  if err != nil {
    return
  }
  events = make(chan string)
  go func(){
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    defer close(events)
    for _ = range ticker.C {
      fi, err := f.Stat(fp)
      if err != nil {
        events <- "file removed"
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
