package fstat

import (
  "os"
  "time"
)

/*
type FileInfo interface {
  Name() string       // base name of the file
  Size() int64        // length in bytes for regular files; system-dependent for others
  Mode() FileMode     // file mode bits
  ModTime() time.Time // modification time
  IsDir() bool        // abbreviation for Mode().IsDir()
  Sys() interface{}   // underlying data source (can return nil)
}
*/

type Finfo struct {
  size int64
}

func (fi Finfo) Name() string { return "" }

func (fi Finfo) Size() int64 {
  return fi.size
}

func (fi Finfo) Mode() os.FileMode {
  return os.FileMode(0)
}

func (fi Finfo) ModTime() time.Time { return time.Now() }

func (fi Finfo) IsDir() bool { return false }

func (fi Finfo) Sys() interface{} {
  return 0
}
