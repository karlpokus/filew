package filew

import (
  "time"
  "fmt"
)

type event struct {
  path, op string
  err error
}

func (ev event) String() string {
  if ev.err != nil {
    return ev.err.Error()
  }
  return fmt.Sprintf("%s %s", ev.path, ev.op)
}

// Watch watches files and returns a channels with file change events
func Watch(root string) (events chan event, err error) {
  base := make(tree)
  err = base.walk(root)
  if err != nil {
    return
  }
  events = make(chan event)
  go func(){
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    defer close(events)
    for _ = range ticker.C { // poll
      fs := make(tree)
      err := fs.walk(root)
      if err != nil {
        events <- event{err: err}
        return
      }
      base.diff(fs, events) // TODO: concurrent op
      base = fs
    }
  }()
  return
}
