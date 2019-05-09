package filew

import (
  "testing"
  "github.com/karlpokus/filew/pkg/fstat"
)

func TestWatch(t *testing.T) {
  mock := &fstat.Mock{}
  events, err := Watch("whatever", mock)
  if err != nil {
    t.Errorf("expected nil, got %s", err)
  }
  mock.EditSize()
  expected := "file changed"
  event := <-events
  if event != expected {
    t.Errorf("expected %s, got %s", expected, event)
  }
}
