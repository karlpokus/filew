package fstat

import (
  "os"
  "sync"
)

type Mock struct {
  sync.Mutex
  Finfo
}

func (m *Mock) Stat(fp string) (os.FileInfo, error) {
  m.Lock()
  defer m.Unlock()
  return m.Finfo, nil
}

func (m *Mock) EditSize() {
  m.Lock()
  defer m.Unlock()
  m.Finfo.size = int64(1)
}
