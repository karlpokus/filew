package mockw

import (
	"errors"

	"github.com/karlpokus/filew/internal/tree"
)

var PathErr = errors.New("oops")

type Mock struct {
	i int
}

func (m *Mock) Walk(root string, t tree.Tree) error {
	defer func() { m.i++ }()
	if m.i == 0 { // base
		t["a"] = int64(1)
		t["b"] = int64(1) // removed
		return nil
	}
	if m.i == 1 { // fs
		t["a"] = int64(2) // updated
		t["c"] = int64(1) // created
		return nil
	}
	return PathErr
}
