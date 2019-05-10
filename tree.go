package filew

import (
  "path/filepath"
  "os"
)

type tree map[string]int64

// walk walks the file system beginning at root and stores all filepaths
// and their size in the tree
func (t tree) walk(root string) error {
  return filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if !fi.IsDir() { // TODO: needs to be regular?
      t[path] = fi.Size()
    }
    return nil
  })
}

// diff reports file events to the events chan
func (base tree) diff(fs tree, events chan event) {
  for fpath, fsSize := range fs {
    baseSize, ok := base[fpath]
    if !ok {
      events <- event{fpath, "created", nil}
      continue
    }
    if fsSize != baseSize {
      events <- event{fpath, "updated", nil}
    }
  }
  for fpath := range base {
    _, ok := fs[fpath]
    if !ok {
      events <- event{fpath, "removed", nil}
    }
  }
}
