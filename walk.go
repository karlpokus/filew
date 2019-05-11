package filew

import (
	"os"
	"path/filepath"

	"github.com/karlpokus/filew/internal/tree"
)

// A Walker stores file paths and sizes in the tree from walking a file system (or similiar)
type Walker interface {
	Walk(string, tree.Tree) error
}

// implements Walker
type fswalk struct{}

func (f fswalk) Walk(root string, t tree.Tree) error {
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

// diff compares trees and reports file events on the events chan
func diff(base, fs tree.Tree, events chan event) {
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
